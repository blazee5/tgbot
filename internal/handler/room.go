package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/blazee5/tgbot/internal/keyboard"
	"github.com/blazee5/tgbot/internal/models"
	"github.com/blazee5/tgbot/pkg/bot_errors"
	"github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
)

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
	fmt.Println(c.Callback().Data)
	go state.Update("room", c.Callback().Data)
	go state.Set(InputTimeState)
	return c.Send("Выберите время", keyboard.BookButtons())
}

func (h *Handler) BookRoom(c tele.Context, state fsm.Context) error {
	defer state.Finish(true)

	var room string

	state.MustGet("room", &room)
	hour := c.Callback().Data

	err := h.service.BookRoom(context.Background(), models.Book{
		UserID: int(c.Sender().ID),
		RoomId: room,
		Time:   hour,
	})

	if errors.Is(err, bot_errors.RoomAlreadyBookedErr) {
		return c.Send("Комната на это время уже занята")
	}

	if err != nil {
		h.log.Infof("error while book room: %v", err)
		return c.Send("Произошла ошибка")
	}

	return c.Send("Комната успешно забронирована")
}
