package main

import (
	"fmt"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/mqtt"
)

func main() {

}

//export start
func _start(rid uint32) int32 {
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
