package handler

import (
	"context"
	"fmt"
	"github.com/blazee5/tgbot/internal/keyboard"
	"github.com/blazee5/tgbot/internal/models"
	"github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
	"strconv"
)

func (h *Handler) Start(c tele.Context, state fsm.Context) error {
	_, err := h.service.User.GetByID(context.Background(), int(c.Sender().ID))

	if err != nil {
		go state.Set(InputFirstName)
		return c.Send("Напишите ваше имя")
	}

	return c.Send("Добро пожаловать", keyboard.MenuButton())
}

func (h *Handler) RegisterUser(c tele.Context, state fsm.Context) error {
	defer state.Finish(true)

	err := h.service.User.CreateUser(context.Background(), models.User{
		ID:        int(c.Sender().ID),
		FirstName: c.Text(),
		LastName:  c.Sender().LastName,
		Username:  c.Sender().Username,
	})

	if err != nil {
		h.log.Infof("error while create user: %v", err)
		return c.Send("error")
	}

	userID := strconv.Itoa(int(c.Sender().ID))

	_, err = h.b.Send(&tele.User{ID: int64(h.cfg.Bot.AdminID)},
		fmt.Sprintf("Новый пользователь @%s", c.Sender().Username), keyboard.VerifyButton(userID))

	if err != nil {
		h.log.Infof("error while send message to admin: %v", err)
		return err
	}

	return c.Send("Успешно")
}

func (h *Handler) VerifyUser(c tele.Context) error {
	userID, err := strconv.Atoi(c.Callback().Data)

	if err != nil {
		h.log.Infof("error while convert userID: %v", err)
		return c.Send("error")
	}

	err = h.service.User.VerifyUser(context.Background(), userID)

	if err != nil {
		h.log.Infof("error while verify user: %v", err)
		return c.Send("Ошибка при подтверждении пользователя")
	}

	return c.Respond(&tele.CallbackResponse{Text: "Пользователь успешно подтвержден"})
}
