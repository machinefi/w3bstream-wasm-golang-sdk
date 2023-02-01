package database

import (
	"encoding/base64"
)

type (
	dBQuery struct {
		Statement string   `json:"statement"`
		Params    []*param `json:"params"`
	}

	param struct {
		Int32   *int32   `json:"int32,omitempty"`
		Int64   *int64   `json:"int64,omitempty"`
		Float32 *float32 `json:"float32,omitempty"`
		Float64 *float64 `json:"float64,omitempty"`
		String  *string  `json:"string,omitempty"`
		Time    *string  `json:"time,omitempty"` //  rfc3339 encoding
		Bool    *bool    `json:"bool,omitempty"`
		Bytes   *string  `json:"bytes,omitempty"` // base64 encoding
	}
)

type sqlInt32 struct {
	data int32
}

func (str *sqlInt32) getParam() *param {
	return &param{
		Int32: &str.data,
	}
}

type sqlInt64 struct {
	data int64
}

func (str *sqlInt64) getParam() *param {
	return &param{
		Int64: &str.data,
	}
}

type sqlFloat32 struct {
	data float32
}

func (str *sqlFloat32) getParam() *param {
	return &param{
		Float32: &str.data,
	}
}

type sqlFloat64 struct {
	data float64
}

func (str *sqlFloat64) getParam() *param {
	return &param{
		Float64: &str.data,
	}
}

type sqlString struct {
	data string
}

func (str *sqlString) getParam() *param {
	return &param{
		String: &str.data,
	}
}

type sqlBool struct {
	data bool
}

func (str *sqlBool) getParam() *param {
	return &param{
		Bool: &str.data,
	}
}

type sqlBytes struct {
	data []byte
}

func (b *sqlBytes) getParam() *param {
	str := base64.StdEncoding.EncodeToString(b.data)
	return &param{
		Bytes: &str,
	}
}

// TODO support time
