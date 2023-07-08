package main

import (
	"bytes"
	"net/http"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/api"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
)

func main() {}

//export start
func _start(rid uint32) int32 {
	req, err := http.NewRequest("GET", "/system/hello", nil)
	if err != nil {
		return -1
	}
	req.Header.Set("name", "w3bstream")

	resp, err := api.Call(req)
	if err != nil {
		return -1
	}

	var buf bytes.Buffer
	if err := resp.Write(&buf); err != nil {
		return -1
	}

	log.Log(string(buf.Bytes()))

	return 0
}
