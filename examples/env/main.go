package main

import (
	"fmt"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/stream"
)

func main() {}

//export start
func _start(rid uint32) int32 {
	key := "env-key"
	val, err := stream.GetEnv(key)
	if err != nil {
		log.Log(fmt.Sprintf("get env from host failed: %v", err))
		return -1
	}
	log.Log(fmt.Sprintf("get env from host: [key:%s] [val:%s]", key, val))
	return 0
}
