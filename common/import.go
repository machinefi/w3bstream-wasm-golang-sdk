package common

//go:wasm-module env
//export ws_log
func WS_log(logLevel, ptr, size uint32) int32

//go:wasm-module env
//export ws_get_data
func WS_get_data(rid, ptr, size uint32) int32

//go:wasm-module env
//export ws_set_data
func WS_set_data(rid, ptr, size uint32) int32

//go:wasm-module env
//export ws_get_db
func WS_get_db(kaddr, ksize, ptr, size uint32) int32

//go:wasm-module env
//export ws_set_db
func WS_set_db(kaddr, ksize, vaddr, vsize uint32) int32

//go:wasm-module env
//export ws_set_sql_db
func WS_set_sql_db(ptr, size uint32) int32

//go:wasm-module env
//export ws_send_tx
func WS_send_tx(kaddr, ksize uint32) (v int32)
