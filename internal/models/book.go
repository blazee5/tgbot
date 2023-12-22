package models

type Book struct {
	UserId   int    `json:"user_id" redis:"user_id"`
	Username string `json:"username" redis:"username"`
	RoomId   string `json:"room_id" redis:"room_id"`
	Time     string `json:"time" redis:"time"`
}
