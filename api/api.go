package api

import (
	"bufio"
	"bytes"
	"errors"
	"net/http"
	"unsafe"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/common"
)

const (
	host   = "w3bstream.com"
	scheme = "w3bstream"
)

func Call(req *http.Request) (*http.Response, error) {
	if req.URL.Host == "" {
		req.URL.Host = host
	}
	if req.URL.Scheme == "" {
		req.URL.Scheme = scheme
	}
	if req.URL.Host != host {
		return nil, errors.New("invalid host")
	}
	if req.URL.Scheme != scheme {
		return nil, errors.New("invalid scheme")
	}
	if req.Header.Get("eventType") == "" {
		return nil, errors.New("missing eventType")
	}

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
