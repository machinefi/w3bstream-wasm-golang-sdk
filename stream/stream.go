package stream

import (
	"fmt"
	"unsafe"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/common"
)

func GetDataByRID(rid uint32) ([]byte, error) {
	addr := uintptr(unsafe.Pointer(new(uint32)))
	size := uintptr(unsafe.Pointer(new(uint32)))

	code := common.WS_get_data(rid, uint32(addr), uint32(size))
	if code != 0 {
		return nil, fmt.Errorf("get data failed: [rid:%d] [code:%d]", rid, code)
	}

	vaddr := *(*uint32)(unsafe.Pointer(addr))
	m := common.Allocations.GetByAddr(vaddr)
	if m == nil {
		return nil, fmt.Errorf("get data by addr failed: [rid:%d] [addr:%d]", rid, vaddr)
	}

	common.Allocations.AddResourceWithMem(rid, m)
	return m.Data, nil
}

func SetDataByRID(rid uint32, data string) {
	addr, size := common.StringToPointer(data)

	_ = common.WS_set_data(rid, addr, size)
}

func GetEnv(key string) (string, error) {
	kaddr, ksize := common.StringToPointer(key)

	vaddr := uintptr(unsafe.Pointer(new(uint32)))
	vsize := uintptr(unsafe.Pointer(new(uint32)))

	code := common.WS_get_env(kaddr, ksize, uint32(vaddr), uint32(vsize))
	if code != 0 {
		return "", fmt.Errorf("get env failed [key:%s] [code:%d]", key, code)
	}

	m := common.Allocations.GetByAddr(*(*uint32)(unsafe.Pointer(vaddr)))
	if m == nil {
		return "", fmt.Errorf("get env failed: [key:%s] [addr:%x]", key, vaddr)
	}

	return string(m.Data), nil
}
