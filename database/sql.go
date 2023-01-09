package database

import (
	"errors"
	"unsafe"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/common"
	"github.com/mailru/easyjson"
)

// Query example:
// To insert a new record could be done like below
// err := database.ExecSQL("INSERT INTO table (ID) VALUES (?);", database.Int32(0))
func ExecSQL(query string, args ...SQLTypes) error {
	params := make([]*param, 0)
	for _, v := range args {
		params = append(params, v.getParam())
	}
	serializedQuery, err := easyjson.Marshal(&dBQuery{
		Statement: query,
		Params:    params,
	})
	if err != nil {
		return errors.New("incorrect sql query format")
	}
	if ret := common.WS_set_sql_db(common.BytesToPointer(serializedQuery)); ret != 0 {
		return errors.New("fail to execute the sql query")
	}
	return nil
}

// Query example:
// To query a new record could be done like below
// result, err := database.QuerySQL("SELECT * FROM table WHERE ID =?;", database.Int32(0))
func QuerySQL(query string, args ...SQLTypes) ([]byte, error) {
	params := make([]*param, 0)
	for _, v := range args {
		params = append(params, v.getParam())
	}
	serializedQuery, err := easyjson.Marshal(&dBQuery{
		Statement: query,
		Params:    params,
	})
	if err != nil {
		return nil, errors.New("incorrect sql query format")
	}

	addr, size := common.BytesToPointer(serializedQuery)
	rAddr := uintptr(unsafe.Pointer(new(uint32)))
	rSize := uintptr(unsafe.Pointer(new(uint32)))
	if ret := common.WS_get_sql_db(addr, size, uint32(rAddr), uint32(rSize)); ret != 0 {
		return nil, errors.New("fail to execute the sql query")
	}

	vaddr := *(*uint32)(unsafe.Pointer(rAddr))
	m := common.Allocations.GetByAddr(vaddr)
	if m == nil {
		return nil, errors.New("fail to get the result of sql query")
	}

	return m.Data, nil
}

type SQLTypes interface {
	getParam() *param
}

func Int32(in int32) SQLTypes {
	return &sqlInt32{
		data: in,
	}
}

func Int64(in int64) SQLTypes {
	return &sqlInt64{
		data: in,
	}
}

func Float32(in float32) SQLTypes {
	return &sqlFloat32{
		data: in,
	}
}

func Float64(in float64) SQLTypes {
	return &sqlFloat64{
		data: in,
	}
}

func String(in string) SQLTypes {
	return &sqlString{
		data: in,
	}
}

func Bool(in bool) SQLTypes {
	return &sqlBool{
		data: in,
	}
}

func Bytes(in []byte) SQLTypes {
	return &sqlBytes{
		data: in,
	}
}

// TODO: support time
