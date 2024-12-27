package database

import "time"


type User struct {
	User_id string `json:"user_id"`
	Photo_id int `json:"photo_id"`
}

type Message struct {
	Message_id string `json:"user_id"`
	Chat_id int
	Status string
	Date time.Time
	Content	string
	Forwarded bool
	Owner string
	Reply int
	Media int
}
type Chat struct {
	Chat_id int
	Group bool
	Name string
	Photo_id int
}