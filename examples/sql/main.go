package main

import (
	"github.com/machinefi/w3bstream-wasm-golang-sdk/database"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
)

func main() {}

//export start
func _start(rid uint32) int32 {
	if err := database.ExecSQL("INSERT INTO table (ID) VALUES (?);", database.Int32(0)); err != nil {
		return -1
	}
	res, err := database.QuerySQL("SELECT * FROM table;")
	if err != nil {
		return -1
	}
	log.Log(string(res))
	return 0
}
