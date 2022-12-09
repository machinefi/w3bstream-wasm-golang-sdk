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
	log.Log(fmt.Sprintf("start received: %d", rid))
	str, err := stream.GetDataByRID(rid)
	if err != nil {
		log.Log("error:" + err.Error())
		return -1
	}

	words := strings.Split(string(str), " ")
	counts := make(map[string]int32)
	for _, w := range words {
		if _, ok := counts[w]; !ok {
			counts[w] = database.GetDB(w) + 1
		} else {
			counts[w]++
		}
	}

	for k, cnt := range counts {
		database.SetDB(k, cnt)
		if _, ok := records[k]; !ok {
			records[k] = cnt
		} else {
			records[k] += cnt
		}
	}
	return 0
}

//export word_count
func _unique(_ uint32) int32 {
	return int32(len(records))
}

var records = make(map[string]int32)
