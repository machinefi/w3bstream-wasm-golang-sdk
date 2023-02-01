package main

import (
	"fmt"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/database"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
)

func main() {}

//export start
func _start(rid uint32) int32 {
	if err := database.Set("key_test", []byte("test")); err != nil {
		return -1
	}
	log.Log("set key success")
	res, err := database.Get("key_test")
	if err != nil {
		log.Log("get key failed")
		return -1
	}
	log.Log("get key success")
	log.Log(fmt.Sprintf("get data %s by key %s", string(res), "key_test"))
	return 0
}
