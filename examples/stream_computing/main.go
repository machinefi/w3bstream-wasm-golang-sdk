package main

import (
	"fmt"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/examples/stream_computing/model"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/stream"
	"github.com/mailru/easyjson"
	"github.com/tidwall/gjson"
)

func main() {

}

//export filterAge
func _filterAge(rid uint32) int32 {
	log.Log(fmt.Sprintf("start rid: %d", rid))

	message, err := stream.GetDataByRID(rid)
	if err != nil {
		log.Log(err.Error())
		return -1
	}

	log.Log(fmt.Sprintf("get start resource all %v: %s", rid, string(message)))

	sourceCustomer := model.SourceCustomer{}
	easyjson.Unmarshal(message, &sourceCustomer)
	age := sourceCustomer.Age
	log.Log(fmt.Sprintf("get start resource age %v: %d", rid, age))

	if age >= 18 {
		log.Log(fmt.Sprintf("filter the Customer's age more than 18 %v: `%s`", rid, string(message)))
		stream.SetDataByRID(rid, "true")
	} else if age < 18 {
		log.Log(fmt.Sprintf("filter the Customer's age less than 18 %v: `%s`", rid, string(message)))
		stream.SetDataByRID(rid, "false")
	}

	return int32(rid)
}

//export mapTax
func _mapTax(rid uint32) int32 {
	log.Log(fmt.Sprintf("mapTax rid: %d", rid))

	message, err := stream.GetDataByRID(rid)
	if err != nil {
		log.Log(err.Error())
		return -1
	}

	log.Log(fmt.Sprintf("get mapTax resource all %v: %s", rid, string(message)))

	sourceCustomer := model.SourceCustomer{}
	easyjson.Unmarshal(message, &sourceCustomer)

	//TODO generate an error
	//common.Log(fmt.Sprintf("get mapTax sourceCustomer %d", sourceCustomer.Age))

	customer := model.Customer{}
	customer.ID = sourceCustomer.ID
	customer.FirstName = sourceCustomer.FirstName
	customer.LastName = sourceCustomer.LastName
	customer.Age = sourceCustomer.Age
	customer.City = sourceCustomer.City

	if customer.Age >= 30 {
		log.Log(fmt.Sprintf("the Customer's age is more than 30 %v: %s", rid, string(message)))
		customer.TaxNumber = "19832106687"
	}

	if b, err := easyjson.Marshal(customer); err != nil {
		log.Log(fmt.Sprintf("%v marshal error", sourceCustomer))
		return -1
	} else {
		stream.SetDataByRID(rid, string(b))
	}

	return int32(rid)
}

//export groupByAge
func _groupByAge(rid uint32) int32 {
	log.Log(fmt.Sprintf("groupByAge rid: %d", rid))

	message, err := stream.GetDataByRID(rid)
	if err != nil {
		log.Log(err.Error())
		return -1
	}

	log.Log(fmt.Sprintf("get groupByAge resource all %v: `%s`", rid, string(message)))

	city := gjson.GetBytes(message, "city").String()
	log.Log(fmt.Sprintf("get groupByAge resource city %v: %s", rid, city))

	stream.SetDataByRID(rid, city)

	return int32(rid)
}
