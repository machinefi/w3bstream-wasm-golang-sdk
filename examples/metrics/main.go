package main

import (
	"fmt"
	"math/rand"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/metrics"
)

func main() {}

//export start
func _start(rid uint32) int32 {
	log.Log(fmt.Sprintf("start rid: %d", rid))

	randomName := []string{
		"MuchSports",
		"Righthaja",
		"BroadcastAut",
		"HunterSkunky",
		"SparklingUnique",
	}
	metricsJSON := `{"user": "` + randomName[rand.Intn(len(randomName))] + `", "score": ` + fmt.Sprintf("%d", rand.Intn(100)) + `}`
	if err := metrics.Submit(metricsJSON); err != nil {
		log.Log(err.Error())
		return -1
	}
	log.Log(fmt.Sprintf("successfully to submit metrics: `%s`", metricsJSON))
	return 0
}
