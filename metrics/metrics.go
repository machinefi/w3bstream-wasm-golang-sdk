//go:build tinygo

package metrics

import (
	"errors"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/common"
)

func Submit(message string) error {
	msgAddr, msgSize := common.StringToPointer(message)
	if ret := common.WS_submit_metrics(msgAddr, msgSize); ret != 0 {
		return errors.New("fail to submit the metrics")
	}
	return nil
}
