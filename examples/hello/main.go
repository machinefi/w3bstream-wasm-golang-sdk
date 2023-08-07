package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/api"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/common"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/stream"
)

func main() {}

//export start
func _start(rid uint32) int32 {
	req, err := http.NewRequest("GET", "/system/hello", nil)
	if err != nil {
		return -1
	}
	req.Header.Set("eventType", "result")
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

//export handle_result
func _handle_result(rid uint32) int32 {
	log.Log(fmt.Sprintf("start rid: %d", rid))

	message, err := stream.GetDataByRID(rid)
	if err != nil {
		log.Log(err.Error())
		return -1
	}

	defer func() {
		if common.FreeResource(rid) {
			log.Log(fmt.Sprintf("resource %v released", rid))
		}
	}()

	log.Log(fmt.Sprintf("get result %v: `%s`", rid, string(message)))
	return 0
}
