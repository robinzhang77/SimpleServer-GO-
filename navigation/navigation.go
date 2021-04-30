package navigation

import (
	"C"
	"fmt"
	"syscall"
)
import (
	"math"
	"time"
	"unsafe"
)

type Navigation struct {
	dll            *syscall.DLL
	callInit       *syscall.Proc
	callLoadMap    *syscall.Proc
	callFreeMap    *syscall.Proc
	callAddAgent   *syscall.Proc
	callGetAnent   *syscall.Proc
	callSetMoveTar *syscall.Proc
	callUpdate     *syscall.Proc
	loopChan       chan interface{}
}

func NewNavigation() *Navigation {
	n := &Navigation{}
	n.dll, _ = syscall.LoadDLL("./navigation/NavigationDLL.dll")
	n.callInit, _ = n.dll.FindProc("recast_init")
	n.callLoadMap, _ = n.dll.FindProc("recast_loadmap")
	n.callFreeMap, _ = n.dll.FindProc("recast_freemap")
	n.callAddAgent, _ = n.dll.FindProc("add_agent")
	n.callGetAnent, _ = n.dll.FindProc("get_agent")
	n.callSetMoveTar, _ = n.dll.FindProc("set_move_target_by_idxs")
	n.callUpdate, _ = n.dll.FindProc("update_tick")

	n.loopChan = make(chan interface{}, 100)
	return n
}

func (n *Navigation) Init() bool {
	ret, _, err := n.callInit.Call()
	if err != nil {
		e := err.(syscall.Errno)
		if e != 0 {
			fmt.Println("init failed : ", err.Error())
			return false
		}
	}
	bVal := (*bool)(unsafe.Pointer(ret))
	if *bVal {
		go n.update()
	}
	return *bVal
}

func (n *Navigation) Fini() {

}

func (n *Navigation) LoadMap(id uint32, strFilePath string) bool {
	fmt.Println("loadmap : ", id, strFilePath)
	char := C.CString(strFilePath)

	r1, _, err := n.callLoadMap.Call(uintptr(id), uintptr(unsafe.Pointer(char)))
	if err != nil {
		e := err.(syscall.Errno)
		if e != 0 {
			fmt.Println("loadmap failed : ", err.Error())
			return false
		}
	}

	bVal := (*bool)(unsafe.Pointer(r1))
	return *bVal
}

func (n *Navigation) FreeMap(id uint32) bool {
	r1, _, err := n.callFreeMap.Call(uintptr(id))
	if err != nil {
		e := err.(syscall.Errno)
		if e != 0 {
			fmt.Println("freemap failed : ", err.Error())
			return false
		}
	}

	bVal := (*bool)(unsafe.Pointer(r1))
	return *bVal
}

func (n *Navigation) AddAgent(id uint32, x, y, z float32, radius float32, speed float32) int {

	pos := [3]float32{x, y, z}
	fmt.Println("add angent :", pos, radius, speed)

	//r1, _, err := n.callAddAgent.Call(uintptr(&id), uintptr(unsafe.Pointer(&pos)), uintptr(unsafe.Pointer(&radius)), uintptr(speed))
	r1, _, err := n.callAddAgent.Call(uintptr(id), uintptr(unsafe.Pointer(&pos)), uintptr(math.Float32bits(radius)), uintptr(math.Float32bits(speed)))
	if err != nil {
		e := err.(syscall.Errno)
		if e != 0 {
			fmt.Println("add agent failed : ", err.Error())
			return -1
		}
	}

	nVal := int(r1)
	return nVal
}

func (n *Navigation) SetMoveTarget(id uint32, idxs []int32, x, y, z float32) {
	nLen := len(idxs)
	pos := [3]float32{x, y, z}
	fmt.Println("set angent move tar :", pos, idxs, nLen)
	_, _, err := n.callSetMoveTar.Call(uintptr(id), uintptr(unsafe.Pointer(&idxs[0])), uintptr(nLen), uintptr(unsafe.Pointer(&pos)))
	if err != nil {
		e := err.(syscall.Errno)
		if e != 0 {
			fmt.Println("set move target failed : ", err.Error())
			return
		}
	}
}

func (n *Navigation) UpdateTick(deltaTime float32) {
	n.callUpdate.Call(uintptr(math.Float32bits(deltaTime)))
}

type AgentInfo struct {
	AgentID int32
	X       float32
	Y       float32
	Z       float32
}

func (n *Navigation) GetAgentsInfo(id uint32, ids []int32) []*AgentInfo {
	agents := []*AgentInfo{}
	nLen := len(ids)
	for i := 0; i < nLen; i++ {
		pos := [3]float32{}
		n.callGetAnent.Call(uintptr(id), uintptr(ids[i]), uintptr(unsafe.Pointer(&pos)))
		//fmt.Println("agent id:", ids[i], " pos:", pos)
		ag := &AgentInfo{}
		ag.AgentID = ids[i]
		ag.X = pos[0]
		ag.Y = pos[1]
		ag.Z = pos[2]

		agents = append(agents, ag)
	}
	return agents
}

func (n *Navigation) update() {
	interval := time.Millisecond * 16
	tiker := time.NewTicker(interval)

	for {
		select {
		case <-n.loopChan:
			//fmt.Println("loopchan11111111111")
		case <-tiker.C:
			n.UpdateTick(0.016)
			//fmt.Println("update...............")
		}
	}
}
