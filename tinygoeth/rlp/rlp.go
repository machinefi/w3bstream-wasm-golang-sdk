package rlp

import (
	"encoding/hex"
	"math/big"
	"strings"
)

// Marshaler is the interface implemented by types that can marshal themselves into valid RLP messages.
type Marshaler interface {
	MarshalRLP() ([]byte, error)
}

// Unmarshaler is the interface implemented by types that can unmarshal a RLP description of themselves
type Unmarshaler interface {
	UnmarshalRLP(buf []byte) error
}

// MarshalRLP marshals an RLP object
func MarshalRLP(m Marshaler) ([]byte, error) {
	return m.MarshalRLP()
}

// UnmarshalRLP unmarshals an RLP object
func UnmarshalRLP(buf []byte, m Unmarshaler) error {
	return m.UnmarshalRLP(buf)
}

// Integer is an interface implemented by all signed integer types
type Integer interface {
	int | int8 | int16 | int32 | int64
}

// Unsigned is an interface implemented by all unsigned integer types
type Unsigned interface {
	uint | uint8 | uint16 | uint32 | uint64
}

// Data is an individual RLP Data element - or an "RLP string"
type Data []byte

// List is a list of RLP elements, which could be either Data or List elements
type List []Element

// Element is an interface implemented by both Data and List elements
type Element interface {
	// When true the Element can safely be cast to List, and when false the Element can safely be cast to Data
	IsList() bool
	// Encode converts the element to a byte array
	Encode() []byte
	//Add adds an element to the list
	Add(v Element) Element
}

// Int converts an integer to an RLP Data element for encoding
func Int[T Integer](i T) Data {
	return BigInt(new(big.Int).SetInt64(int64(i)))
}

// Uint converts an unsigned integer to an RLP Data element for encoding
func Uint[T Unsigned](i T) Data {
	return BigInt(new(big.Int).SetUint64(uint64(i)))
}

// BigInt converts a big.Int to an RLP Data element for encoding
func BigInt(i *big.Int) Data {
	return Data(i.Bytes())
}

// String converts a string to an RLP Data element for encoding
func String(s string) Data {
	return Bytes([]byte(s))
}

// Bytes converts a byte array to an RLP Data element for encoding
func Bytes(b []byte) Data {
	return Data(b)
}

// Bool converts a boolean to an RLP Data element for encoding
func Bool(b bool) Data {
	if b {
		return Data{1}
	}
	return Data{}
}

// NewList creates a new empty RLP List element
func NewList() Element {
	return List{}
}

// Hex converts a hex encoded string (with or without 0x prefix) to an RLP Data element for encoding
func Hex(s string) (Data, error) {
	b, err := hex.DecodeString(strings.TrimPrefix(s, "0x"))
	if err != nil {
		return nil, err
	}
	return Data(b), nil
}

// MustHex panics if hex decoding fails
func MustHex(s string) Data {
	b, err := Hex(s)
	if err != nil {
		panic(err)
	}
	return b
}

// Int is a convenience function to convert the bytes within an RLP Data element to an integer (big endian encoding)
func (r Data) Int() *big.Int {
	if r == nil {
		return nil
	}
	i := new(big.Int)
	return i.SetBytes(r)
}

// Encode encodes this individual RLP Data element
func (r Data) Encode() []byte {
	return encodeBytes(r, false)
}

// Add adds an element to the list
func (r Data) Add(v Element) Element {
	panic("cannot set on data")
}

// IsList is false for individual RLP Data elements
func (r Data) IsList() bool {
	return false
}

// Bool is a convenience function to convert the bytes within an RLP Data element to a boolean
func (r Data) Bool() bool {
	return len(r) > 0 && r[0] != 0
}

// String is a convenience function to convert the bytes within an RLP Data element to a string
func (r Data) String() string {
	return string(r)
}

// Bytes is a convenience function to convert the bytes within an RLP Data element to a byte array
func (r Data) Bytes() []byte {
	return []byte(r)
}

//Byte is a convenience function to convert the bytes within an RLP Data element to a byte
func (r Data) Byte() byte {
	return r[0]
}

// Uint is a convenience function to convert the bytes within an RLP Data element to an unsigned integer (big endian encoding)
func (r Data) Uint() uint {
	return uint(r.Int().Uint64())
}

// Hex is a convenience function to convert the bytes within an RLP Data element to a hex encoded string
func (r Data) Hex() string {
	return hex.EncodeToString(r)
}

// Encode encodes the RLP List to a byte array, including recursing into child arrays
func (l List) Encode() []byte {
	if len(l) == 0 {
		return encodeBytes([]byte{}, true)
	}
	var concatenation []byte
	for _, entry := range l {
		concatenation = append(concatenation, entry.Encode()...)
	}
	return encodeBytes(concatenation, true)

}

// IsList returns true for list elements
func (l List) IsList() bool {
	return true
}

// Add adds an element to the list
func (l List) Add(v Element) Element {
	return append(l, v)
}
