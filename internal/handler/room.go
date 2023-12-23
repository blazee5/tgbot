package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/blazee5/tgbot/internal/keyboard"
	"github.com/blazee5/tgbot/pkg/bot_errors"
	"github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
	"strings"
)

func (h *Handler) GetBooks(c tele.Context) error {
	books, err := h.service.Room.GetBooksByID(context.Background(), int(c.Sender().ID))

	if err != nil {
		return c.Send("error")
	}

	if len(books) == 0 {
		return c.Send("У вас пока что нет броней")
	}

	for _, book := range books {
		text := fmt.Sprintf(`
Комната №%s
Время %s:00
`, book.RoomId, book.Time)
		err = c.Send(text, keyboard.CancelBookButton(book.RoomId, book.Time))

		if err != nil {
			h.log.Infof("error while get books: %v", err)
			return err
		}
	}

	return nil
}

func (h *Handler) SelectRoom(c tele.Context) error {
	user, err := h.service.User.GetByID(context.Background(), int(c.Sender().ID))

	if err != nil {
		h.log.Infof("error while book room: %v", err)
		return c.Send("error")
	}

	if !user.IsVerified {
		return c.Send("Ваш аккаунт не подтвержден, ожидайте")
	}

	return c.Send("Выберите комнату", keyboard.RoomsButton())
}

func (h *Handler) SelectTime(c tele.Context, state fsm.Context) error {
	roomID := c.Callback().Data

	go state.Update("room", roomID)
	go state.Set(InputTimeState)

	err := c.Delete()

	if err != nil {
		return err
	}

	timeBtns, err := keyboard.BookButtons(roomID, h.service.Room)

	if err != nil {
		h.log.Infof("error while get time buttons: %v", err)
		return err
	}

	if len(timeBtns.InlineKeyboard[0]) == 0 {
		return c.Send("На сегодня вся комната занята(")
	}

	return c.Send("Выберите время", timeBtns)
}

func (h *Handler) BookRoom(c tele.Context, state fsm.Context) error {
	defer state.Finish(true)

	var room string

	state.MustGet("room", &room)
	hour := c.Callback().Data

	username, err := h.service.BookRoom(context.Background(), int(c.Sender().ID), room, hour)

	if errors.Is(err, bot_errors.RoomAlreadyBookedErr) {
		return c.Send(fmt.Sprintf("Комната на это время уже занята (%s)", username))
	}

	if err != nil {
		h.log.Infof("error while book room: %v", err)
		return c.Send("Произошла ошибка")
	}

	err = c.Delete()

	if err != nil {
		return err
	}

	return c.Send("Комната успешно забронирована")
}

func (h *Handler) CancelBook(c tele.Context) error {
	data := strings.Split(c.Callback().Data, ":")
	roomId := data[0]
	hour := data[1]

	err := h.service.Room.CancelBook(context.Background(), roomId, hour, int(c.Sender().ID))

	if err != nil {
		h.log.Infof("error while cancel book: %v", err)
		return c.Send("Произошла ошибка")
	}

	err = c.Delete()

	if err != nil {
		return err
	}

	return c.Send("Бронь успешно отменена")
}
