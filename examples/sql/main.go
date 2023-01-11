//go:build tinygo

package main

import (
	"github.com/machinefi/w3bstream-wasm-golang-sdk/database"
)

func main() {}

//export start
func _start(rid uint32) int32 {
	if err := database.ExecSQL("INSERT INTO table (ID) VALUES (?);", database.Int32(0)); err != nil {
		return -1
	}
	return 0
}
