package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/api"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/common"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/stream"
)

func main() {}

//export start
func _start(rid uint32) int32 {
	data := `{"chainID": 4690,"hash": "fcaf377ff3cc785d60c58de7e121d6a2e79e1c58c189ea8641f3ea61f7605285"}`

	req, err := http.NewRequest("GET", "/system/read_tx", strings.NewReader(data))
	if err != nil {
		return -1
	}
	req.Header.Set("eventType", "result")

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
