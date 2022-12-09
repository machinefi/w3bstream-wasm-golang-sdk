package main

import (
	"fmt"

	"github.com/mailru/easyjson"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/examples/easyjson/model"
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
	log.Log("wasm received: " + string(message))
	student := model.Student{}
	easyjson.Unmarshal(message, &student)
	log.Log("wasm get struct.name from json:" + student.Name)
	log.Log("wasm change student name to Jane ")
	student.Name = "Jane"
	log.Log("wasm get new name from the struct:" + student.Name)
	msg, err := easyjson.Marshal(student)
	log.Log("wasm get json from struct: " + string(msg))
	return 0
}
