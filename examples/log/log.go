//go:build tinygo

package main

import (
	"fmt"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/common"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/stream"
)

// main is required for TinyGo to compile to Wasm.
func main() {}

//export start
func _start(rid uint32) int32 {
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

	log.Log(fmt.Sprintf("get resource %v: `%s`", rid, string(message)))
	return 0
}
