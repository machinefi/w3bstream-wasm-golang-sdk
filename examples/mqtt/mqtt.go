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

//export start
func _fetchAndPush(rid uint32) int32 {
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

	topic, payload = "test1", []byte("123")
	log.Log(fmt.Sprintf("send [%s %s]", topic, string(payload)))
	if err = mqtt.SendMqttMsg(topic, payload); err != nil {
		log.Log(fmt.Sprintf("pub message failed: [topic: %s] [err: %v]", topic, err.Error()))
		return -1
	}
	log.Log(fmt.Sprintf("pub message: [topic: %s] [msg: %v]", topic, string(payload)))

	topic = "test2"
	log.Log(fmt.Sprintf("send [%s nil]", topic))
	err = mqtt.SendMqttMsg(topic, nil)
	if err != nil {
		log.Log(fmt.Sprintf("pub message failed: [topic: %s] [err: %v]", topic, err.Error()))
		return -1
	}
	log.Log(fmt.Sprintf("pub empty message: [topic: %s]", topic))
	return 0
}
