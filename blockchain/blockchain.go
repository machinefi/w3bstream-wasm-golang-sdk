//go:build tinygo

package blockchain

import "github.com/machinefi/w3bstream-wasm-golang-sdk/common"

func SendTx(key string) int32 {
	addr, size := common.StringToPointer(key)
	return common.WS_send_tx(addr, size)
}
