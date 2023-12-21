package models

type Book struct {
	UserID int    `json:"user_id" redis:"user_id"`
	RoomId string `json:"room_id" redis:"room_id"`
	Time   string `json:"time" redis:"time"`
}
