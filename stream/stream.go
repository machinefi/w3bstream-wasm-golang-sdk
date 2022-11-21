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
