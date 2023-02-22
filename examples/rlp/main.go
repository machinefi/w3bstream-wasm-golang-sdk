package main

import (
	"fmt"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/tinygoeth/rlp"
)

var errInvalidRlp = fmt.Errorf("invalid rlp")

type simple struct {
	Data1 []byte
	Data2 [][]byte
	Data3 uint64
}

// MarshalRLP implements rlp.Marshaler, and serializes s into the Ethereum RLP format.
// equivalent to go-ethereum: rlp.EncodeToBytes(s)
func (s *simple) MarshalRLP() ([]byte, error) {
	data2 := rlp.NewList()
	for _, d := range s.Data2 {
		data2 = data2.Add(rlp.Data(d))
	}
	a := rlp.List{
		rlp.Data(s.Data1),
		data2,
		rlp.Int(s.Data3),
	}.Encode()
	return a, nil
}

// UnmarshalRLP implements rlp.Unmarshaler, and loads the s fields from a RLP message.
func (s *simple) UnmarshalRLP(buf []byte) error {
	elems, pos, err := rlp.Decode(buf)
	if err != nil {
		return err
	}
	getElem := func() rlp.Element {
		if elems.IsList() {
			elem := elems.(rlp.List)[0]
			elems = elems.(rlp.List)[1:]
			return elem
		}
		return elems
	}
	if pos != len(buf) {
		return errInvalidRlp
	}
	if elems.IsList() {
		s.Data1 = []byte(getElem().(rlp.Data))
		ele := getElem()
		if !ele.IsList() {
			return errInvalidRlp
		}
		for _, e := range ele.(rlp.List) {
			if e.IsList() {
				return errInvalidRlp
			}
			s.Data2 = append(s.Data2, []byte(e.(rlp.Data)))
		}
		s.Data3 = uint64(getElem().(rlp.Data).Uint())
	} else {
		return errInvalidRlp
	}
	return nil
}

func main() {}

//export start
func _start(rid uint32) int32 {
	s := &simple{
		Data1: []byte("hello world"),
		Data2: [][]byte{
			[]byte("hello"),
			[]byte("world"),
		},
		Data3: 12345,
	}
	a, err := rlp.MarshalRLP(s)
	log.Log(fmt.Sprintf("marshal success, encoded: %x", a))
	if err != nil {
		log.Log("marshal rlp failed" + err.Error())
		return -1
	}
	s2 := &simple{}
	err = rlp.UnmarshalRLP(a, s2)
	if err != nil {
		log.Log("unmarshal rlp failed" + err.Error())
		return -1
	}
	log.Log(fmt.Sprintf("unmarshal success, s2.Data1=%s s2.Data2 length=%d s2.Data2[0]=%s s2.Data2[1]=%s s2.Date3=%d",
		s2.Data1, len(s2.Data2), s2.Data2[0], s2.Data2[1], s2.Data3))
	return 0
}
