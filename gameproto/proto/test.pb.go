// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/test.proto

package gameproto

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgID int32

const (
	MsgID_NONE           MsgID = 0
	MsgID_NotifyPlrID    MsgID = 1
	MsgID_ReqCreateAgent MsgID = 2
	MsgID_ReqMove        MsgID = 3
)

var MsgID_name = map[int32]string{
	0: "NONE",
	1: "NotifyPlrID",
	2: "ReqCreateAgent",
	3: "ReqMove",
}

var MsgID_value = map[string]int32{
	"NONE":           0,
	"NotifyPlrID":    1,
	"ReqCreateAgent": 2,
	"ReqMove":        3,
}

func (x MsgID) String() string {
	return proto.EnumName(MsgID_name, int32(x))
}

func (MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fc70169f84046e97, []int{0}
}

//登陆成功，服务器分配一个唯一ID
type NotifyPlrUUID struct {
	Plr string `protobuf:"bytes,1,opt,name=plr,proto3" json:"plr,omitempty"`
}

func (m *NotifyPlrUUID) Reset()         { *m = NotifyPlrUUID{} }
func (m *NotifyPlrUUID) String() string { return proto.CompactTextString(m) }
func (*NotifyPlrUUID) ProtoMessage()    {}
func (*NotifyPlrUUID) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc70169f84046e97, []int{0}
}
func (m *NotifyPlrUUID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NotifyPlrUUID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NotifyPlrUUID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NotifyPlrUUID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyPlrUUID.Merge(m, src)
}
func (m *NotifyPlrUUID) XXX_Size() int {
	return m.Size()
}
func (m *NotifyPlrUUID) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyPlrUUID.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyPlrUUID proto.InternalMessageInfo

func (m *NotifyPlrUUID) GetPlr() string {
	if m != nil {
		return m.Plr
	}
	return ""
}

//造兵请求
type CreateAgentReq struct {
	Plr    string  `protobuf:"bytes,1,opt,name=plr,proto3" json:"plr,omitempty"`
	Id     int32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	X      float32 `protobuf:"fixed32,3,opt,name=x,proto3" json:"x,omitempty"`
	Y      float32 `protobuf:"fixed32,4,opt,name=y,proto3" json:"y,omitempty"`
	Z      float32 `protobuf:"fixed32,5,opt,name=z,proto3" json:"z,omitempty"`
	Radius float32 `protobuf:"fixed32,6,opt,name=radius,proto3" json:"radius,omitempty"`
	Speed  float32 `protobuf:"fixed32,7,opt,name=speed,proto3" json:"speed,omitempty"`
}

func (m *CreateAgentReq) Reset()         { *m = CreateAgentReq{} }
func (m *CreateAgentReq) String() string { return proto.CompactTextString(m) }
func (*CreateAgentReq) ProtoMessage()    {}
func (*CreateAgentReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc70169f84046e97, []int{1}
}
func (m *CreateAgentReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateAgentReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateAgentReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateAgentReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAgentReq.Merge(m, src)
}
func (m *CreateAgentReq) XXX_Size() int {
	return m.Size()
}
func (m *CreateAgentReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAgentReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAgentReq proto.InternalMessageInfo

func (m *CreateAgentReq) GetPlr() string {
	if m != nil {
		return m.Plr
	}
	return ""
}

func (m *CreateAgentReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateAgentReq) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *CreateAgentReq) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *CreateAgentReq) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func (m *CreateAgentReq) GetRadius() float32 {
	if m != nil {
		return m.Radius
	}
	return 0
}

func (m *CreateAgentReq) GetSpeed() float32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

//造兵返回（广播）
type CreateAgentNotify struct {
	Plr     string  `protobuf:"bytes,1,opt,name=plr,proto3" json:"plr,omitempty"`
	Id      int32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	AgentId int32   `protobuf:"varint,3,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	X       float32 `protobuf:"fixed32,4,opt,name=x,proto3" json:"x,omitempty"`
	Y       float32 `protobuf:"fixed32,5,opt,name=y,proto3" json:"y,omitempty"`
	Z       float32 `protobuf:"fixed32,6,opt,name=z,proto3" json:"z,omitempty"`
}

func (m *CreateAgentNotify) Reset()         { *m = CreateAgentNotify{} }
func (m *CreateAgentNotify) String() string { return proto.CompactTextString(m) }
func (*CreateAgentNotify) ProtoMessage()    {}
func (*CreateAgentNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc70169f84046e97, []int{2}
}
func (m *CreateAgentNotify) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateAgentNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateAgentNotify.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateAgentNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAgentNotify.Merge(m, src)
}
func (m *CreateAgentNotify) XXX_Size() int {
	return m.Size()
}
func (m *CreateAgentNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAgentNotify.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAgentNotify proto.InternalMessageInfo

func (m *CreateAgentNotify) GetPlr() string {
	if m != nil {
		return m.Plr
	}
	return ""
}

func (m *CreateAgentNotify) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateAgentNotify) GetAgentId() int32 {
	if m != nil {
		return m.AgentId
	}
	return 0
}

func (m *CreateAgentNotify) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *CreateAgentNotify) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *CreateAgentNotify) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

//请求移动
type SetMoveReq struct {
	Plr string  `protobuf:"bytes,1,opt,name=plr,proto3" json:"plr,omitempty"`
	Ids []int32 `protobuf:"varint,2,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	X   float32 `protobuf:"fixed32,3,opt,name=x,proto3" json:"x,omitempty"`
	Y   float32 `protobuf:"fixed32,4,opt,name=y,proto3" json:"y,omitempty"`
	Z   float32 `protobuf:"fixed32,5,opt,name=z,proto3" json:"z,omitempty"`
}

func (m *SetMoveReq) Reset()         { *m = SetMoveReq{} }
func (m *SetMoveReq) String() string { return proto.CompactTextString(m) }
func (*SetMoveReq) ProtoMessage()    {}
func (*SetMoveReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc70169f84046e97, []int{3}
}
func (m *SetMoveReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetMoveReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetMoveReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetMoveReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetMoveReq.Merge(m, src)
}
func (m *SetMoveReq) XXX_Size() int {
	return m.Size()
}
func (m *SetMoveReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SetMoveReq.DiscardUnknown(m)
}

var xxx_messageInfo_SetMoveReq proto.InternalMessageInfo

func (m *SetMoveReq) GetPlr() string {
	if m != nil {
		return m.Plr
	}
	return ""
}

func (m *SetMoveReq) GetIds() []int32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *SetMoveReq) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *SetMoveReq) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *SetMoveReq) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func init() {
	proto.RegisterEnum("gameproto.MsgID", MsgID_name, MsgID_value)
	proto.RegisterType((*NotifyPlrUUID)(nil), "gameproto.NotifyPlrUUID")
	proto.RegisterType((*CreateAgentReq)(nil), "gameproto.CreateAgentReq")
	proto.RegisterType((*CreateAgentNotify)(nil), "gameproto.CreateAgentNotify")
	proto.RegisterType((*SetMoveReq)(nil), "gameproto.SetMoveReq")
}

func init() { proto.RegisterFile("proto/test.proto", fileDescriptor_fc70169f84046e97) }

var fileDescriptor_fc70169f84046e97 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0xe3, 0xa4, 0x4e, 0xdb, 0xd3, 0x7b, 0x8b, 0xb1, 0x10, 0x32, 0x4b, 0x14, 0x32, 0x45,
	0x0c, 0x30, 0xf0, 0x04, 0xd0, 0x32, 0x64, 0x68, 0x41, 0x46, 0x5d, 0x58, 0x50, 0x50, 0x0e, 0x95,
	0xa5, 0x42, 0xda, 0xc4, 0xa0, 0xb6, 0x8f, 0xc0, 0xc4, 0x63, 0x31, 0x76, 0x64, 0x44, 0xed, 0x8b,
	0x20, 0x3b, 0x15, 0x6a, 0x05, 0x43, 0xb7, 0xff, 0x3b, 0xb2, 0xa5, 0xef, 0x9c, 0x1f, 0xd8, 0xb8,
	0xc8, 0x75, 0x7e, 0xa6, 0xb1, 0xd4, 0xa7, 0x36, 0xf2, 0xe6, 0x30, 0x7d, 0x42, 0x1b, 0xa3, 0x63,
	0xf8, 0xdf, 0xcf, 0xb5, 0x7a, 0x9c, 0xdd, 0x8c, 0x8a, 0xc1, 0x20, 0xe9, 0x72, 0x06, 0xde, 0x78,
	0x54, 0x08, 0x12, 0x92, 0xb8, 0x29, 0x4d, 0x8c, 0xde, 0x08, 0xb4, 0x3b, 0x05, 0xa6, 0x1a, 0x2f,
	0x86, 0xf8, 0xac, 0x25, 0x4e, 0x7e, 0x3f, 0xe2, 0x6d, 0x70, 0x55, 0x26, 0xdc, 0x90, 0xc4, 0x54,
	0xba, 0x2a, 0xe3, 0xff, 0x80, 0x4c, 0x85, 0x17, 0x92, 0xd8, 0x95, 0x64, 0x6a, 0x68, 0x26, 0x6a,
	0x15, 0xcd, 0x0c, 0xcd, 0x05, 0xad, 0x68, 0xce, 0x0f, 0xc1, 0x2f, 0xd2, 0x4c, 0xbd, 0x94, 0xc2,
	0xb7, 0xa3, 0x35, 0xf1, 0x03, 0xa0, 0xe5, 0x18, 0x31, 0x13, 0x75, 0x3b, 0xae, 0x20, 0x9a, 0xc3,
	0xfe, 0x86, 0x4b, 0xa5, 0xbe, 0x83, 0xce, 0x11, 0x34, 0x52, 0xf3, 0xe1, 0x5e, 0x65, 0xd6, 0x8a,
	0xca, 0xba, 0xe5, 0x64, 0x6d, 0x5a, 0xdb, 0x32, 0xa5, 0x5b, 0xa6, 0xfe, 0xda, 0x34, 0xba, 0x03,
	0xb8, 0x45, 0xdd, 0xcb, 0x5f, 0xf1, 0xef, 0x1b, 0x30, 0xf0, 0x54, 0x56, 0x0a, 0x37, 0xf4, 0x62,
	0x2a, 0x4d, 0xdc, 0xfd, 0x0a, 0x27, 0x1d, 0xa0, 0xbd, 0x72, 0x98, 0x74, 0x79, 0x03, 0x6a, 0xfd,
	0xeb, 0xfe, 0x15, 0x73, 0xf8, 0x1e, 0xb4, 0x7e, 0xaa, 0x49, 0xba, 0x8c, 0x70, 0x0e, 0x6d, 0x89,
	0x93, 0x8d, 0xf5, 0x99, 0xcb, 0x5b, 0x50, 0x97, 0x38, 0x31, 0x4e, 0xcc, 0xbb, 0x14, 0x1f, 0xcb,
	0x80, 0x2c, 0x96, 0x01, 0xf9, 0x5a, 0x06, 0xe4, 0x7d, 0x15, 0x38, 0x8b, 0x55, 0xe0, 0x7c, 0xae,
	0x02, 0xe7, 0xc1, 0xb7, 0x6d, 0x9f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xba, 0xf5, 0xaa,
	0x0c, 0x02, 0x00, 0x00,
}

func (m *NotifyPlrUUID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NotifyPlrUUID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NotifyPlrUUID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Plr) > 0 {
		i -= len(m.Plr)
		copy(dAtA[i:], m.Plr)
		i = encodeVarintTest(dAtA, i, uint64(len(m.Plr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CreateAgentReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateAgentReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateAgentReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Speed != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Speed))))
		i--
		dAtA[i] = 0x3d
	}
	if m.Radius != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Radius))))
		i--
		dAtA[i] = 0x35
	}
	if m.Z != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Z))))
		i--
		dAtA[i] = 0x2d
	}
	if m.Y != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Y))))
		i--
		dAtA[i] = 0x25
	}
	if m.X != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.X))))
		i--
		dAtA[i] = 0x1d
	}
	if m.Id != 0 {
		i = encodeVarintTest(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Plr) > 0 {
		i -= len(m.Plr)
		copy(dAtA[i:], m.Plr)
		i = encodeVarintTest(dAtA, i, uint64(len(m.Plr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CreateAgentNotify) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateAgentNotify) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateAgentNotify) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Z != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Z))))
		i--
		dAtA[i] = 0x35
	}
	if m.Y != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Y))))
		i--
		dAtA[i] = 0x2d
	}
	if m.X != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.X))))
		i--
		dAtA[i] = 0x25
	}
	if m.AgentId != 0 {
		i = encodeVarintTest(dAtA, i, uint64(m.AgentId))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintTest(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Plr) > 0 {
		i -= len(m.Plr)
		copy(dAtA[i:], m.Plr)
		i = encodeVarintTest(dAtA, i, uint64(len(m.Plr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SetMoveReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetMoveReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetMoveReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Z != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Z))))
		i--
		dAtA[i] = 0x2d
	}
	if m.Y != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Y))))
		i--
		dAtA[i] = 0x25
	}
	if m.X != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.X))))
		i--
		dAtA[i] = 0x1d
	}
	if len(m.Ids) > 0 {
		dAtA2 := make([]byte, len(m.Ids)*10)
		var j1 int
		for _, num1 := range m.Ids {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintTest(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Plr) > 0 {
		i -= len(m.Plr)
		copy(dAtA[i:], m.Plr)
		i = encodeVarintTest(dAtA, i, uint64(len(m.Plr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTest(dAtA []byte, offset int, v uint64) int {
	offset -= sovTest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NotifyPlrUUID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Plr)
	if l > 0 {
		n += 1 + l + sovTest(uint64(l))
	}
	return n
}

func (m *CreateAgentReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Plr)
	if l > 0 {
		n += 1 + l + sovTest(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTest(uint64(m.Id))
	}
	if m.X != 0 {
		n += 5
	}
	if m.Y != 0 {
		n += 5
	}
	if m.Z != 0 {
		n += 5
	}
	if m.Radius != 0 {
		n += 5
	}
	if m.Speed != 0 {
		n += 5
	}
	return n
}

func (m *CreateAgentNotify) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Plr)
	if l > 0 {
		n += 1 + l + sovTest(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTest(uint64(m.Id))
	}
	if m.AgentId != 0 {
		n += 1 + sovTest(uint64(m.AgentId))
	}
	if m.X != 0 {
		n += 5
	}
	if m.Y != 0 {
		n += 5
	}
	if m.Z != 0 {
		n += 5
	}
	return n
}

func (m *SetMoveReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Plr)
	if l > 0 {
		n += 1 + l + sovTest(uint64(l))
	}
	if len(m.Ids) > 0 {
		l = 0
		for _, e := range m.Ids {
			l += sovTest(uint64(e))
		}
		n += 1 + sovTest(uint64(l)) + l
	}
	if m.X != 0 {
		n += 5
	}
	if m.Y != 0 {
		n += 5
	}
	if m.Z != 0 {
		n += 5
	}
	return n
}

func sovTest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTest(x uint64) (n int) {
	return sovTest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NotifyPlrUUID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NotifyPlrUUID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NotifyPlrUUID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Plr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreateAgentReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateAgentReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateAgentReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Plr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.X = float32(math.Float32frombits(v))
		case 4:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Y = float32(math.Float32frombits(v))
		case 5:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Z", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Z = float32(math.Float32frombits(v))
		case 6:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Radius", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Radius = float32(math.Float32frombits(v))
		case 7:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Speed", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Speed = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreateAgentNotify) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateAgentNotify: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateAgentNotify: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Plr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AgentId", wireType)
			}
			m.AgentId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AgentId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.X = float32(math.Float32frombits(v))
		case 5:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Y = float32(math.Float32frombits(v))
		case 6:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Z", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Z = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SetMoveReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetMoveReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetMoveReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Plr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTest
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Ids = append(m.Ids, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTest
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTest
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTest
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Ids) == 0 {
					m.Ids = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTest
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Ids = append(m.Ids, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Ids", wireType)
			}
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.X = float32(math.Float32frombits(v))
		case 4:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Y = float32(math.Float32frombits(v))
		case 5:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Z", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Z = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTest
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTest
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTest
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTest = fmt.Errorf("proto: unexpected end of group")
)