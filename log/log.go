//go:build tinygo

package log

import "github.com/machinefi/w3bstream-wasm-golang-sdk/common"

func Log(message string) {
	ptr, size := common.StringToPointer(message)
	_ = common.WS_log(4, ptr, size) // logInfoLevel = 4 logWarnLevel = 3 logErrorLevel = 2
}
