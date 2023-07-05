//go:build tinygo

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
//export ws_function_call
func WS_function_call(kaddr, ksize, ptr, size uint32) int32

//go:wasm-module env
//export ws_set_db
func WS_set_db(kaddr, ksize, vaddr, vsize uint32) int32

//go:wasm-module env
//export ws_set_sql_db
func WS_set_sql_db(ptr, size uint32) int32

//go:wasm-module env
//export ws_get_sql_db
func WS_get_sql_db(ptr, size uint32, vaddr, vsize uint32) int32

//go:wasm-module env
//export ws_send_tx
func WS_send_tx(chainID uint32, kaddr, ksize uint32, vaddr, vsize uint32) (v int32)

//go:wasm-module env
//export ws_call_contract
func WS_call_contract(chainID uint32, kaddr, ksize uint32, vaddr, vsize uint32) (v int32)

//go:wasm-module env
//export ws_get_env
func WS_get_env(kaddr, ksize, vaddr, vsize uint32) int32

//go:wasm-module env
//export ws_get_mqtt_msg
func WS_get_mqtt_msg(rid, topicaddr, topicsize, pladdr, plsize uint32) int32

//go:wasm-module env
//export ws_send_mqtt_msg
func WS_send_mqtt_msg(topicAddr, topicSize, msgAddr, msgSize uint32) int32
