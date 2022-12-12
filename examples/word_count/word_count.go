package main

import (
	"fmt"
	"strings"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/database"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/stream"
)

func main() {}

//export start
func _start(rid uint32) int32 {
	log.Log(fmt.Sprintf("start rid: %d", rid))
	str, err := stream.GetDataByRID(rid)
	if err != nil {
		log.Log("error: " + err.Error())
		return -1
	}

	words := strings.Split(string(str), " ")
	records := make(map[string]int32)
	for _, w := range words {
		if _, ok := records[w]; !ok {
			records[w] = database.GetDB(w) + 1
		} else {
			records[w]++
		}
	}

	for k, cnt := range records {
		database.SetDB(k, cnt)
	}
	return 0
}
