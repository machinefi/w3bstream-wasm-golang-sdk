package blockchain

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"unsafe"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/common"
)

func SendTx(chainID uint32, to string, value *big.Int, data string) (string, error) {
	str := fmt.Sprintf(fmt.Sprintf(
		`{
			"to": "%s",
			"value": "%s",
			"data": "%s"
		}`,
		to,
		value.String(),
		strings.TrimPrefix(data, "0x"),
	))
	addr, size := common.StringToPointer(str)

	vaddr := uintptr(unsafe.Pointer(new(uint32)))
	vsize := uintptr(unsafe.Pointer(new(uint32)))
	if ret := common.WS_send_tx(chainID, addr, size, uint32(vaddr), uint32(vsize)); ret != 0 {
		return "", errors.New("sent tx failed")
	}

	m := common.Allocations.GetByAddr(*(*uint32)(unsafe.Pointer(vaddr)))
	if m == nil {
		return "", errors.New("sent tx failed")
	}
	return string(m.Data), nil
}

func SendTxWithPrivateKey(chainID uint32, to string, value *big.Int, data, privateKey string) (string, error) {
	str := fmt.Sprintf(fmt.Sprintf(
		`{
			"to": "%s",
			"value": "%s",
			"data": "%s",
			"privateKey": "%s"
		}`,
		to,
		value.String(),
		strings.TrimPrefix(data, "0x"),
		strings.TrimPrefix(privateKey, "0x"),
	))
	addr, size := common.StringToPointer(str)

	vaddr := uintptr(unsafe.Pointer(new(uint32)))
	vsize := uintptr(unsafe.Pointer(new(uint32)))
	if ret := common.WS_send_tx_with_private_key(chainID, addr, size, uint32(vaddr), uint32(vsize)); ret != 0 {
		return "", errors.New("sent tx with private key failed")
	}

	m := common.Allocations.GetByAddr(*(*uint32)(unsafe.Pointer(vaddr)))
	if m == nil {
		return "", errors.New("sent tx with private key failed")
	}
	return string(m.Data), nil
}

func CallContract(chainID uint32, to string, data string) ([]byte, error) {
	str := fmt.Sprintf(fmt.Sprintf(
		`{
			"to": "%s",
			"data": "%s"
		}`,
		to,
		strings.TrimPrefix(data, "0x"),
	))
	addr, size := common.StringToPointer(str)

	vaddr := uintptr(unsafe.Pointer(new(uint32)))
	vsize := uintptr(unsafe.Pointer(new(uint32)))
	if ret := common.WS_call_contract(chainID, addr, size, uint32(vaddr), uint32(vsize)); ret != 0 {
		return nil, errors.New("sent tx failed")
	}

	m := common.Allocations.GetByAddr(*(*uint32)(unsafe.Pointer(vaddr)))
	if m == nil {
		return nil, errors.New("sent tx failed")
	}
	return m.Data, nil
}
