package keyboard

import (
	"context"
	"fmt"
	"github.com/blazee5/tgbot/internal/service"
	tele "gopkg.in/telebot.v3"
	"strconv"
	"time"
)

var (
	menu          = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnBookRooms  = menu.Text("Забронировать")
	BtnGetBooks   = menu.Text("Мои брони")
	rooms         = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnRoomOne    = rooms.Data("Комната 1", "select-room", "1")
	BtnRoomTwo    = rooms.Data("Комната 2", "select-room", "2")
	verify        = &tele.ReplyMarkup{}
	BtnVerify     = verify.Data("Подтвердить", "user-verify")
	book          = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnTime       = book.Data("", "book-room")
	BtnCancelBook = book.Data("Отменить", "book-cancel")
	BtnBack       = book.Data("Назад", "back")
)

func MenuButton() *tele.ReplyMarkup {
	menu.Reply(
		menu.Row(BtnBookRooms, BtnGetBooks),
	)

	return menu
}

func RoomsButton() *tele.ReplyMarkup {
	rooms.Inline(
		rooms.Row(BtnRoomOne, BtnRoomTwo),
	)

	return rooms
}

func VerifyButton(id string) *tele.ReplyMarkup {
	BtnVerify.Data = id

	verify.Inline(verify.Row(BtnVerify))

	return verify
}

func BookButtons(roomId string, service service.Room) (*tele.ReplyMarkup, error) {
	currentHour := time.Now().Hour()

	if currentHour < 10 {
		currentHour = 10
	}

	btns := make([]tele.Btn, 0)

	for i := currentHour + 1; i < 20; i++ {
		username, err := service.GetBook(context.Background(), roomId, strconv.Itoa(i))

		if err != nil {
			return nil, err
		}

		if username != "" {
			continue
		}

		BtnTime.Text = fmt.Sprintf("%02d:00", i)
		BtnTime.Data = strconv.Itoa(i)
		btns = append(btns, BtnTime)
	}
	book.Inline(book.Row(btns...), book.Row(BtnBack))

	return book, nil
}

func CancelBookButton(roomId, time string) *tele.ReplyMarkup {
	BtnCancelBook.Data = fmt.Sprintf("%s:%s", roomId, time)

	book.Inline(book.Row(BtnCancelBook))

	return book
}
