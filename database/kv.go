package database

import (
	"bytes"
	"encoding/binary"
	"unsafe"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/common"
)

func GetDB(key string) int32 {
	addr, size := common.StringToPointer(key)

	rAddr := uintptr(unsafe.Pointer(new(uint32)))
	rSize := uintptr(unsafe.Pointer(new(uint32)))

	if ret := common.WS_get_db(addr, size, uint32(rAddr), uint32(rSize)); ret != 0 {
		return 0
	}

	vaddr := *(*uint32)(unsafe.Pointer(rAddr))
	m := common.Allocations.GetByAddr(vaddr)
	if m == nil {
		return 0
	}

	var ret int32
	buf := bytes.NewBuffer(m.Data)
	binary.Read(buf, binary.LittleEndian, &ret)
	return ret
}

func SetDB(key string, v int32) {
	addr, size := common.StringToPointer(key)

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, v)
	vaddr, vsize := common.BytesToPointer(buf.Bytes())

	// TODO: error handle
	_ = common.WS_set_db(addr, size, vaddr, vsize)
}
