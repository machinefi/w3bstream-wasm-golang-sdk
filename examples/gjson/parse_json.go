//go:build tinygo

package main

import (
	"fmt"

	"github.com/tidwall/gjson"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/stream"
)

func main() {}

//export start
func _start(rid uint32) int32 {
	log.Log(fmt.Sprintf("start received: %d", rid))
	message, err := stream.GetDataByRID(rid)
	if err != nil {
		log.Log("error: " + err.Error())
		return -1
	}
	res := string(message)
	log.Log("wasm received message: " + res)
	log.Log("wasm get name(json string) from json: " + gjson.Get(res, "name").String())
	log.Log("wasm get name.age(int) from json: " + gjson.Get(res, "name.age").String())
	log.Log("wasm get friends(array) from json: " + gjson.Get(res, "friends").String())
	log.Log("wasm get friends[0].nets(array) from json: " + gjson.Get(res, "friends.0.nets").String())
	return 0
}
