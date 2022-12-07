package database_types

type (
	dBQuery struct {
		Statement string   `json:"statement"`
		Params    []*param `json:"params"`
	}

	param struct {
		Int32   int32   `json:"int32,omitempty"`
		Int64   int64   `json:"int64,omitempty"`
		Float32 float32 `json:"float32,omitempty"`
		Float64 float64 `json:"float64,omitempty"`
		String  string  `json:"string,omitempty"`
		Time    string  `json:"time,omitempty"` //  rfc3339 encoding
		Bool    bool    `json:"bool,omitempty"`
		Bytes   string  `json:"bytes,omitempty"` // base64 encoding
	}
)
