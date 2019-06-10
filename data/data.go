package data

import "time"

//easyjson:json
type Todo struct {
	Id          int       `json:"Id"`
	Name        string    `json:"Name"`
	Description string    `json:"Desciption"`
	Completed   bool      `json:"Completed"`
	UserId      int       `json:"UserId"`
	StartTime   time.Time `json:"StartTime"`
}

//easyjson:json
type User struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

//easyjson:json
type Error struct {
	Errorname string `json:"Errorname"`
}

//easyjson:json
type TodoArray []Todo

//easyjson:json
type UserArray []User
