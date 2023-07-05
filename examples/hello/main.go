package main

import (
	"bytes"
	gohttp "net/http"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/http"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
)

func main() {}

//export start
func _start(rid uint32) int32 {
	req, err := gohttp.NewRequest("GET", "/system/hello", nil)
	if err != nil {
		return -1
	}
	req.Header.Set("name", "w3bstream")

	resp, err := http.Do(req)
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
