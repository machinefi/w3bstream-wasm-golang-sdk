package main

import (
	"fmt"
	// "strconv"

	"github.com/mailru/easyjson"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/database"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/examples/tesla_poc/model"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/stream"
)

func main() {}

//export start
func _start(rid uint32) int32 {
	log.Log(fmt.Sprintf("start received: %d", rid))
	message, err := stream.GetDataByRID(rid)
	if err != nil {
		log.Log("error: " + err.Error())
		return -1
	}
	log.Log("wasm received: " + string(message))
	tesla := model.Tesla{}
	easyjson.Unmarshal(message, &tesla)
	log.Log("wasm get car name from the struct"+tesla.Name)
	
	// newerr := database.ExecSQL("INSERT INTO wasm_project__tesla_poc20.teslastream_data (c_carname,c_latitude,c_longitude,c_percent_remaining,c_range,c_capacity,c_plugged_in,c_charge_state,c_odometer,c_back_left,c_back_right,c_front_left,c_front_right) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?);", database.String(tesla.Name),database.Float64(tesla.Latitude),database.Float64(tesla.Longitude),database.Float64(tesla.Percent_remaining),database.Float64(tesla.Range),database.Float64(tesla.Capacity),database.Bool(tesla.Plugged_in),database.String(tesla.Charge_state),database.Float64(tesla.Odometer),database.Float64(tesla.Back_left),database.Float64(tesla.Back_right),database.Float64(tesla.Front_left),database.Float64(tesla.Front_right))
	// newerr := database.ExecSQL("INSERT INTO wasm_project__tesla_poc20.teslastream_data (c_carname,c_latitude,c_longitude,c_percent_remaining,c_range,c_capacity,c_plugged_in,c_charge_state,c_odometer,c_back_left,c_back_right,c_front_left,c_front_right) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?);", database.String(tesla.Name),database.Float64(tesla.Latitude),database.Float64(tesla.Longitude),database.Float64(tesla.Percent_remaining),database.Float64(tesla.Range),database.Float64(tesla.Capacity),database.Bool(true),database.String(tesla.Charge_state),database.Float64(tesla.Odometer),database.Float64(tesla.Back_left),database.Float64(tesla.Back_right),database.Float64(tesla.Front_left),database.Float64(tesla.Front_right))
	newerr := database.ExecSQL("INSERT INTO wasm_project__tesla_poc20.teslastream_data (c_carname,c_latitude,c_longitude,c_percent_remaining,c_range,c_capacity,c_plugged_in,c_charge_state,c_odometer,c_back_left,c_back_right,c_front_left,c_front_right) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?);", database.String(tesla.Name),database.Float64(tesla.Latitude),database.Float64(tesla.Longitude),database.Float64(tesla.Percent_remaining),database.Float64(tesla.Range),database.Float64(tesla.Capacity),database.Bool(true),database.String(tesla.Charge_state),database.Float64(tesla.Odometer),database.Float64(tesla.Back_left),database.Float64(tesla.Back_right),database.Float64(tesla.Front_left),database.Float64(tesla.Front_right))


	if newerr != nil {
		log.Log("error: " + newerr.Error())
		return -1
	}
	log.Log("finished!")

	// msg, err := easyjson.Marshal(tesla)
	// log.Log("wasm get json from struct: " + string(msg))
	return 0
}
