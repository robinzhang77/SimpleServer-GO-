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

type navigation struct {
	dll            *syscall.DLL
	callInit       *syscall.Proc
	callLoadMap    *syscall.Proc
	callFreeMap    *syscall.Proc
	callAddAgent   *syscall.Proc
	callSetMoveTar *syscall.Proc
	callUpdate     *syscall.Proc
}

func NewNavigation() *navigation {
	n := &navigation{}
	n.dll, _ = syscall.LoadDLL("./navigation/NavigationDLL.dll")
	n.callInit, _ = n.dll.FindProc("recast_init")
	n.callLoadMap, _ = n.dll.FindProc("recast_loadmap")
	n.callFreeMap, _ = n.dll.FindProc("recast_freemap")
	n.callAddAgent, _ = n.dll.FindProc("add_agent")
	n.callSetMoveTar, _ = n.dll.FindProc("set_move_target_by_idxs")
	n.callUpdate, _ = n.dll.FindProc("update_tick")

	return n
}

func (n *navigation) Init() bool {
	ret, _, err := n.callInit.Call()
	if err != nil {
		e := err.(syscall.Errno)
		if e != 0 {
			fmt.Println("init failed : ", err.Error())
			return false
		}
	}
	bVal := (*bool)(unsafe.Pointer(ret))
	return *bVal
}

func (n *navigation) Fini() {

}

func (n *navigation) LoadMap(id uint32, strFilePath string) bool {
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

func (n *navigation) FreeMap(id uint32) bool {
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

func (n *navigation) AddAgent(id uint32, x, y, z float32, radius float32, speed float32) int {

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

func (n *navigation) SetMoveTarget(id uint32, idxs []uint32, x, y, z float32) {
	nLen := len(idxs)
	pos := [3]float32{x, y, z}
	fmt.Println("set angent move tar :", pos, idxs)
	_, _, err := n.callSetMoveTar.Call(uintptr(id), uintptr(unsafe.Pointer(&idxs)), uintptr(nLen), uintptr(unsafe.Pointer(&pos)))
	if err != nil {
		e := err.(syscall.Errno)
		if e != 0 {
			fmt.Println("set move target failed : ", err.Error())
			return
		}
	}
}

func init() {
	fmt.Println("navigation init")

	go func() {
		for {
			fmt.Println("vavigation update tick")
			time.Sleep(time.Second)
		}
	}()
}
