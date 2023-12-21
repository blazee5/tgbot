package handler

import (
	"context"
	"fmt"
	"github.com/blazee5/tgbot/internal/keyboard"
	"github.com/blazee5/tgbot/internal/models"
	tele "gopkg.in/telebot.v3"
	"strconv"
)

func (h *Handler) Hello(c tele.Context) error {
	id, err := h.service.User.CreateUser(context.Background(), models.User{
		ID:        int(c.Sender().ID),
		FirstName: c.Sender().FirstName,
		LastName:  c.Sender().LastName,
		Username:  c.Sender().Username,
	})

	if err != nil {
		h.log.Infof("error while create user: %v", err)
		return c.Send("Произошла ошибка, попробуйте снова")
	}
	userID := strconv.Itoa(int(c.Sender().ID))

	if id != 0 {
		_, err = h.b.Send(&tele.User{ID: int64(h.cfg.Bot.AdminID)},
			fmt.Sprintf("Новый пользователь @%s", c.Sender().Username), keyboard.VerifyButton(userID))

		if err != nil {
			h.log.Infof("error while send message to admin: %v", err)
			return nil
		}
	}

	return c.Send("hello", keyboard.MenuButton())
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
