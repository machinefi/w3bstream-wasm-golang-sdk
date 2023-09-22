package main

import (
	"bytes"
	"fmt"
	"io"
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
	imageID := "3145991386, 3471678490, 3632776032, 2595288688, 1478438623, 4259749138, 987879707, 1456846509" //Range
	privateInput := "16"
	publicInput := "4,30"
	// Stark or Snark
	receiptType := "Stark"
	data := fmt.Sprintf(`{"imageID": "%s","privateInput": "%s","publicInput": "%s", "receiptType": "%s"}`,
		imageID, privateInput, publicInput, receiptType)
	log.Log(data)

	req, err := http.NewRequest("POST", "/system/gen_zk_proof", strings.NewReader(data))
	if err != nil {
		log.Log(err.Error())
		return -1
	}
	req.Header.Set("eventType", "result")

	resp, err := api.Call(req)
	if err != nil {
		log.Log(err.Error())
		return -1
	}

	var buf bytes.Buffer
	if err := resp.Write(&buf); err != nil {
		log.Log(err.Error())
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

	resp, err := api.ConvResponse(message)
	if err != nil {
		log.Log(err.Error())
		return -1
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Log(err.Error())
		return -1
	}

	log.Log(fmt.Sprintf("get result: %v, status: %v, information: %v", rid, resp.Status, string(body)))
	return 0
}
