package model

import "time"

//easyjson:json
type Tesla struct {
	Id       int        `json:"id"`
	Name     string     `json:"tesla_name"`
	Latitude float64    `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Percent_remaining  float64   `json:"percent_remaining"`
	Range    float64    `json:"range"`
	Capacity float64    `json:"capacity"`
	Plugged_in  bool  `json:"is_plugged_in"`
	Charge_state  string `json:"state"`
	Odometer  float64   `json:"odometer"`
	Back_left  float64  `json:"back_left"`
	Back_right  float64  `json:"back_right"`
	Front_left  float64  `json:"front_left"`
	Front_right  float64  `json:"front_right"`
	Time     time.Time  `json:"update_time"`
}