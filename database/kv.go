package database

import (
	"errors"
	"unsafe"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/common"
)

func Get(key string) ([]byte, error) {
	addr, size := common.StringToPointer(key)

	rAddr := uintptr(unsafe.Pointer(new(uint32)))
	rSize := uintptr(unsafe.Pointer(new(uint32)))

	if ret := common.WS_get_db(addr, size, uint32(rAddr), uint32(rSize)); ret != 0 {
		return nil, errors.New("fail to get the data from db")
	}

	vaddr := *(*uint32)(unsafe.Pointer(rAddr))
	m := common.Allocations.GetByAddr(vaddr)
	if m == nil {
		return nil, errors.New("fail to get the data from db")
	}

	return m.Data, nil
}

func Set(key string, value []byte) error {
	addr, size := common.StringToPointer(key)
	vaddr, vsize := common.BytesToPointer(value)

	if ret := common.WS_set_db(addr, size, vaddr, vsize); ret != 0 {
		return errors.New("fail to set the data into db")
	}
	return nil
}
