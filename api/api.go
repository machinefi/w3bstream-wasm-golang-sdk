package api

import (
	"errors"
	"io"
	"net/http"
	"unsafe"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/api/model"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/common"
	"github.com/mailru/easyjson"
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
		return nil, errors.New("missing event type")
	}

	var body []byte
	if req.Body != nil {
		var err error
		body, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
	}
	mr := model.HttpRequest{
		Method: req.Method,
		Url:    req.URL.String(),
		Header: req.Header,
		Body:   body,
	}
	buf, err := easyjson.Marshal(&mr)
	if err != nil {
		return nil, err
	}

	addr, size := common.BytesToPointer(buf)

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

	return ConvResponse(m.Data)
}
