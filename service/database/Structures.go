package database

import "time"

type User struct {
	User_id  string `json:"user_id"`
	Photo_id int    `json:"photo_id"`
}

type Message struct {
	Message_id string `json:"message_id"`
	Chat_id    int    `json:"chat_id"`
	Status     int    `json:"status"`
	Content    string `json:"content"`
	Forwarded  int    `json:"forwarded"`
	Owner      string `json:"user_id"`
	Reply      int    `json:"reply"`
	Media      int    `json:"photo_id"`
	Date       string `json:"date"`
}

type ChangeUserId struct {
	User_id   string `json:"old_user_id"`
	User_id_2 string `json:"new_user_id"`
}

type UserChatCombo struct {
	User_id string `json:"user_id"`
	Chat_id int    `json:"chat_id"`
}

type Reaction struct {
	Message_id string `json:"message_id"`
	Owner      string `json:"user_id"`
	Reaction   string `json:"reaction"`
}

type ForwardedMessage struct {
	Message_id string    `json:"message_id"`
	Chat_id    int       `json:"first_chat_id"`
	Chat_id_2  int       `json:"second_chat_id"`
	Status     string    `json:"status"`
	Date       time.Time `json:"date"`
	Content    string    `json:"content"`
	Forwarded  bool      `json:"forwarded"`
	Owner      string    `json:"user_id"`
	Reply      int       `json:"reply"`
	Media      int       `json:"photo_id"`
}

type Chat struct {
	Chat_id    int `json:"first_chat_id"`
	Chat_group bool
	Chat_photo int
	Chat_name  string
}
