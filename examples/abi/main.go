package main

import (
	"bytes"
	"encoding/hex"
	"math/big"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/tinygoeth/abi"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/tinygoeth/address"
)

func main() {}

//export start
func _start(rid uint32) int32 {
	m, err := abi.NewMethod("confirm(string,address,uint32,bytes,uint256,uint32)")
	if err != nil {
		log.Log("abi new error: " + err.Error())
		return -1
	}
	data, err := m.Pack("AAAAA", address.HexToAddress("0xbb6c0bdfec9b5f77050dd9b73d645aa0aa08f20d"), uint32(1), []byte("BBBBB"), big.NewInt(2), uint32(3))
	if err != nil {
		log.Log("pack error: " + err.Error())
		return -1
	}
	log.Log("pack result: " + hex.EncodeToString(data))

	want, _ := hex.DecodeString("2409556e00000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000bb6c0bdfec9b5f77050dd9b73d645aa0aa08f20d00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000005414141414100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000054242424242000000000000000000000000000000000000000000000000000000")
	if !bytes.Equal(data, want) {
		log.Log("pack error: got " + hex.EncodeToString(data) + ", want " + hex.EncodeToString(want))
	}

	return 0
}
