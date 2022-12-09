package model

import "time"

//easyjson:json
type Student struct {
	Id       int       `json:"id"`
	Name     string    `json:"student_name"`
	School   School    `json:"student_school"`
	Birthday time.Time `json:"birthday"`
}
