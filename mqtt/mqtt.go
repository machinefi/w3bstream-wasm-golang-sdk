//go:build tinygo

package mqtt

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/common"
)

func GetMqttMsg(rid uint32) (string, []byte, error) {
	topicaddr := uintptr(unsafe.Pointer(new(uint32)))
	topicsize := uintptr(unsafe.Pointer(new(uint32)))
	pladdr := uintptr(unsafe.Pointer(new(uint32)))
	plsize := uintptr(unsafe.Pointer(new(uint32)))

	code := common.WS_get_mqtt_msg(rid, uint32(topicaddr), uint32(topicsize), uint32(pladdr), uint32(plsize))
	if code != 0 {
		return "", nil, fmt.Errorf("get mqtt msg failed: [rid:%d] [code:%d]", rid, code)
	}

	addr := *(*uint32)(unsafe.Pointer(topicaddr))
	memtopic := common.Allocations.GetByAddr(addr)
	if memtopic == nil {
		return "", nil, fmt.Errorf("get topic failed: [rid:%d] [topic:%d]", rid, addr)
	}
	topic := string(memtopic.Data)

	addr = *(*uint32)(unsafe.Pointer(pladdr))
	mempl := common.Allocations.GetByAddr(addr)
	if mempl == nil {
		return topic, nil, nil
	}
	payload := mempl.Data

	return topic, payload, nil
}

func SendMqttMsg(topic string, payload []byte) error {
	topicAddr, topicSize := common.StringToPointer(topic)
	msgAddr, msgSize := common.BytesToPointer(payload)

	if ret := common.WS_send_mqtt_msg(topicAddr, topicSize, msgAddr, msgSize); ret != 0 {
		return errors.New("fail to send message to mqtt")
	}
	return nil
}
