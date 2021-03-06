package main

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/golang/protobuf/proto"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	s := newServer()
	s.Start()
}

const PacketHeaderSize = 4                   // 数据包头部的大小
const DefaultMaxPacketSize = 4 * 1024 * 1024 // 默认最大包大小:10M

//////////////////////////////////////////////////////////
var (
	ErrType           = errors.New("error type")
	ErrPacketTooLarge = errors.New("packet too large")
	ErrPacketNil      = errors.New("packet is nil")
	ErrLogicPanic     = errors.New("logic goroutine panic")
	ErrClosed         = errors.New("connection has been closed")
)

type marshaler struct {
}

func NewMarshaler() *marshaler {
	return new(marshaler)
}

func (c *marshaler) Marshal(obj interface{}) ([]byte, error) {
	pm, ok := obj.(proto.Message)
	if !ok {
		return nil, ErrType
	}

	return proto.Marshal(pm)
}

func (c *marshaler) Unmarshal(data []byte, obj interface{}) error {
	pm, ok := obj.(proto.Message)
	if !ok {
		return ErrType
	}

	return proto.Unmarshal(data, pm)
}

func (c *marshaler) String() string {
	return "pb"
}

////////////////////////////////////////////////////

// 消息包
type Packet struct {
	Len  uint32
	Body []byte
}

type Server struct {
	listener    *net.TCPListener
	connections *sync.Map
	closed      uint32
}

type Connection struct {
	closeHandler func(*Connection) // 连接关闭回调
	conn         *net.TCPConn      // TCP连接
	readChan     chan *Packet      // 入包通道
	writeChan    chan *Packet      // 出包通道
	pool         sync.Pool         // Packet内存池优化，降低GC
	connClosed   uint32            // 标记连接是否已断开
	codec        *codecImpl
	closeErr     error
	ctx          context.Context
	closeFn      func() // 主动关闭函数
}

// 构造新包
func NewPacket() *Packet {
	return &Packet{}
}

// 判断是否为空包
func (p *Packet) IsEmpty() bool {
	return p.Len == 0 && len(p.Body) == 0
}

// 清空
func (p *Packet) Clear() {
	p.Len = 0
	p.Body = p.Body[:0]
}

// 克隆包
func (p *Packet) Clone() *Packet {
	newPacket := &Packet{}
	newPacket.Len = p.Len
	newPacket.Body = make([]byte, len(p.Body))
	copy(newPacket.Body, p.Body)
	return newPacket
}

//创建一个server
func newServer() *Server {
	return &Server{closed: 0,
		connections: new(sync.Map)}
}

//创建一个连接
func newConnection(c *net.TCPConn) *Connection {
	var readChanCap = 10
	var writeChanCap = 10

	tcpConn := &Connection{
		conn:      c,
		readChan:  make(chan *Packet, readChanCap),
		writeChan: make(chan *Packet, writeChanCap),
		codec:     NewCodec(binary.BigEndian),
	}

	tcpConn.pool.New = func() interface{} {
		return NewPacket()
	}
	tcpConn.ctx, tcpConn.closeFn = context.WithCancel(context.Background())

	return tcpConn
}

//连接关闭
func (c *Connection) Close(n uint32) {
	fmt.Println("connect closed")
}

func (c *Connection) registerCloseHandler(h func(*Connection)) {
	c.closeHandler = h
}

func (s *Server) Start() {
	fmt.Println("Server start , listen 10124")
	if err := s.listen(); err != nil {
		fmt.Println("server listen error", err)
		return
	}

	var tempDelay time.Duration // how long to sleep on accept failure

	//go func() {
	for {
		conn, err := s.listener.AcceptTCP()
		if err != nil {
			if opErr, ok := err.(*net.OpError); ok && opErr.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2

					if max := 1 * time.Second; tempDelay > max {
						tempDelay = max
					}
				}
				time.Sleep(tempDelay)
				continue
			}

			if s.isClosed() {
				return
			}

			return
		}

		tempDelay = 0

		if s.isClosed() {
			return
		}

		s.onConnect(conn)
	}
	//}()
}

func (s *Server) listen() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:10124")
	if err != nil {
		return err
	}

	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}

	s.listener = ln

	return nil
}

func (s *Server) isClosed() bool {
	return atomic.LoadUint32(&s.closed) == 1
}

// Stop server
func (s *Server) Stop() {
	if !atomic.CompareAndSwapUint32(&s.closed, 0, 1) {
		return
	}

	_ = s.listener.Close()

	s.connections.Range(func(k, v interface{}) bool {
		c := v.(Connection)
		c.Close(0)
		return true
	})
}

func (s *Server) onConnect(conn *net.TCPConn) {

	fmt.Println("create new connection")
	c := newConnection(conn)

	c.registerCloseHandler(s.onClose)

	s.connections.Store(conn.RemoteAddr(), c)

	c.loop()
}

func (s *Server) onClose(c *Connection) {
	_, ok := s.connections.Load(c.conn.RemoteAddr())
	if !ok {
		fmt.Println("OnClose connection has been closed from ", c.conn.RemoteAddr())
		return
	}

	s.connections.Delete(c.conn.RemoteAddr())
	fmt.Println("removed connection")
}

//处理连接的网络数据
func (c *Connection) loop() {
	loopers := []func(){c.readLoop, c.writeLoop, c.handleLoop}
	for _, l := range loopers {
		looper := l
		go func() {
			looper()
		}()
	}
}

// 被动关闭
func (c *Connection) onClose() {
	if atomic.CompareAndSwapUint32(&c.connClosed, 0, 1) {
		_ = c.conn.Close()
	}

	if c.closeHandler != nil {
		c.closeHandler(c)
	}
}

// 获得内存池中的Packet
func (c *Connection) GetPoolPacket() *Packet {
	packet := c.pool.Get().(*Packet)
	packet.Clear()
	return packet
}

// 释放Packet到内存池中
func (c *Connection) PutPoolPacket(packet *Packet) {
	c.pool.Put(packet)
}

type codecImpl struct {
	byteOrder binary.ByteOrder
}

// NewCodec
func NewCodec(byteOrder binary.ByteOrder) *codecImpl {
	c := &codecImpl{byteOrder: byteOrder}

	return c
}

// 读取数据
func (c *codecImpl) Read(reader io.Reader, p *Packet) error {

	buf := make([]byte, PacketHeaderSize)

	if _, err := io.ReadFull(reader, buf); err != nil {
		return err
	}

	p.Len = c.byteOrder.Uint32(buf[0:4])
	// p.ID = c.byteOrder.Uint32(buf[4:8])
	// p.Flag = c.byteOrder.Uint32(buf[8:12])
	// p.Cmd = c.byteOrder.Uint32(buf[12:16])
	// p.Ec = c.byteOrder.Uint32(buf[16:20])

	// 判断Body是否过大
	if p.Len >= DefaultMaxPacketSize {
		return ErrPacketTooLarge
	}

	if p.Body == nil || cap(p.Body) < int(p.Len) {
		// 当Body的容量不足时，进行扩容
		p.Body = make([]byte, p.Len)
	} else {
		p.Body = p.Body[:int(p.Len)]
	}

	if _, err := io.ReadFull(reader, p.Body); err != nil {
		return err
	}

	return nil
}

// 写入数据
func (c *codecImpl) Write(writer io.Writer, p *Packet) error {
	if p == nil {
		return ErrPacketNil
	}

	// 判断Body是否过大
	if len(p.Body) >= DefaultMaxPacketSize {
		return ErrPacketTooLarge
	}

	p.Len = uint32(len(p.Body))

	buf := make([]byte, PacketHeaderSize)

	c.byteOrder.PutUint32(buf[0:4], p.Len)
	// c.byteOrder.PutUint32(buf[4:8], p.ID)
	// c.byteOrder.PutUint32(buf[8:12], p.Flag)
	// c.byteOrder.PutUint32(buf[12:16], p.Cmd)
	// c.byteOrder.PutUint32(buf[16:20], p.Ec)

	if _, err := writer.Write(buf); err != nil {
		return err
	}

	_, err := writer.Write(p.Body)
	return err
}

var (
	ErrReadTimeout = errors.New("read time out")
	ErrEOF         = errors.New("EOF")
	ErrNormal      = errors.New("normal error")
	ErrSvrStop     = errors.New("server stop")
)

// 是否已关闭
func (c *Connection) IsClosed() bool {
	return atomic.LoadUint32(&c.connClosed) == 1
}

func (c *Connection) readLoop() {
	defer c.onClose()

	for {

		var packet = c.GetPoolPacket()

		if err := c.codec.Read(c.conn, packet); err != nil {
			c.closeErr = ErrNormal
			emptyPacket := packet.IsEmpty()
			c.PutPoolPacket(packet)

			// 被动关闭
			if err == io.EOF {
				c.closeErr = ErrEOF
				return
			}

			// 如果连接已经关闭，则返回
			if c.IsClosed() {
				return
			}

			// 网络错误处理
			if netErr, ok := err.(net.Error); ok {
				// 如果超时了，则关闭连接
				if netErr.Timeout() {
					c.closeErr = ErrReadTimeout
					return
				}

				if netErr.Temporary() {
					continue
				}
			}

			// 负载均衡检测，直接关闭连接
			if emptyPacket {
				return
			}

			return
		}

		c.readChan <- packet
	}
}

func (c *Connection) writeLoop() {
	for {
		select {
		case p := <-c.writeChan:
			packet := p
			//exception.RecoveryFunc(func() {
			err := c.write(packet)
			if err != nil {
				if !c.IsClosed() {
					fmt.Println("Send packet failed: ", err)
				}

				return
			}
			//})
			c.PutPoolPacket(packet)

		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Connection) flush() {
	for {
		select {
		case p := <-c.writeChan:
			packet := p
			//exception.RecoveryFunc(func() {
			err := c.write(packet)
			if err != nil {
				if !c.IsClosed() {
					fmt.Println("flush packet failed: ", err)
				}

				return
			}
			//})
			c.PutPoolPacket(packet)
		default:
			return
		}
	}
}

func (c *Connection) write(packet *Packet) (err error) {
	err = c.codec.Write(c.conn, packet)
	if err != nil {
		return err
	}

	return nil
}

// 发送消息
func (c *Connection) Send(msg *Packet) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = ErrLogicPanic
		}
	}()

	select {
	case <-c.ctx.Done():
		return ErrClosed
	default:
		c.writeChan <- msg
	}

	return
}

func (c *Connection) handleLoop() {
	for {
		select {
		case p := <-c.readChan:

			packet := p
			//exception.RecoveryFunc(func() {
			resp := c.Process(packet)
			if resp == nil {
				return
			}

			err := c.Send(resp)

			if err != nil && !c.IsClosed() {
				fmt.Println("handle send failed:", packet.Len, err)
			}
			//})

			c.PutPoolPacket(packet)

		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Connection) Process(msg *Packet) *Packet {
	var ret *Packet
	// switch msg.Cmd {
	// case CmdMsgPlayerEnterFight.Value():
	// 	ret = h.PlayerEnter(conn, msg.Body)
	// }
	return ret
}
