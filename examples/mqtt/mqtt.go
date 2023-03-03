//go:build tinygo

package main

import (
	"fmt"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/common"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/mqtt"
)

// main is required for TinyGo to compile to Wasm.
func main() {}

//export fetch
func _fetch(rid uint32) int32 {
	log.Log(fmt.Sprintf("start rid: %d", rid))

	topic, payload, err := mqtt.GetMqttMsg(rid)
	if err != nil {
		log.Log(err.Error())
		return -1
	}

	defer func() {
		if common.FreeResource(rid) {
			log.Log(fmt.Sprintf("resource %v released", rid))
		}
	}()

	log.Log(fmt.Sprintf("get mqtt [rid: %d] [topic: %s] [pl: %s]", rid, topic, string(payload)))
	return 0
}

//export push
func _push() {
	log.Log("start to send message to mqtt topic[mqtt_test]")
	topic, msg := "mqtt_test", `{"key":"w3bstream mqtt test"}`
	err := mqtt.SendMqttMsg(topic, msg)
	if err != nil {
		log.Log("send message to mqtt failed")
		return -1
	}
	log.Log(fmt.Sprintf("send [%s] to mqtt topic [%s]", msg, topic))
	return 0
}
