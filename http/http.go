package http

import (
	"bufio"
	"bytes"
	"errors"
	"net/http"
	"unsafe"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/common"
)

func Do(req *http.Request) (*http.Response, error) {
	req.URL.Host = "w3bstream.com"
	req.URL.Scheme = "w3bstream"

	var buf bytes.Buffer

	if err := req.Write(&buf); err != nil {
		return nil, err
	}

	addr, size := common.BytesToPointer(buf.Bytes())

	rAddr := uintptr(unsafe.Pointer(new(uint32)))
	rSize := uintptr(unsafe.Pointer(new(uint32)))

	if ret := common.WS_api_call(addr, size, uint32(rAddr), uint32(rSize)); ret != 0 {
		return nil, errors.New("fail to get the data from host")
	}

	vaddr := *(*uint32)(unsafe.Pointer(rAddr))
	m := common.Allocations.GetByAddr(vaddr)
	if m == nil {
		return nil, errors.New("fail to get the data from host")
	}

	return http.ReadResponse(bufio.NewReader(bytes.NewBuffer(m.Data)), nil)
}
