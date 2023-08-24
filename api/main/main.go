package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/api/model"
	"github.com/mailru/easyjson"
)

func main() {
	data := `{"chainID": 80001,"hash": "0x7b221cd72ccf6ef65b8d05e50807c2c3dcee984cb5f6c916e7484a3e871ef017"}`

	req, err := http.NewRequest("GET", "/system/read_tx", strings.NewReader(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("eventType", "result")

	body, error := io.ReadAll(req.Body)
	if error != nil {
		panic(err)
	}
	rr := model.HttpRequest{
		Url:    req.URL.String(),
		Header: req.Header,
		Body:   body,
	}

	msg, err := easyjson.Marshal(&rr)
	fmt.Println(string(msg))

	nr := model.HttpRequest{}
	if err := easyjson.Unmarshal(msg, &nr); err != nil {
		panic(err)
	}

	fmt.Println(nr)
}
