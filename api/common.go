package api

import (
	"bytes"
	"io"
	"net/http"

	"github.com/mailru/easyjson"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/api/model"
)

func ConvResponse(data []byte) (*http.Response, error) {
	nr := model.HttpResponse{}
	if err := easyjson.Unmarshal(data, &nr); err != nil {
		return nil, err
	}

	return &http.Response{
		Status:     nr.Status,
		StatusCode: nr.StatusCode,
		Proto:      nr.Proto,
		Header:     nr.Header,
		Body:       io.NopCloser(bytes.NewReader(nr.Body)),
	}, nil
}
