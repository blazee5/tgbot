package handler

import (
	"github.com/blazee5/tgbot/config"
	"github.com/blazee5/tgbot/internal/keyboard"
	"github.com/blazee5/tgbot/internal/service"
	"github.com/vitaliy-ukiru/fsm-telebot"
	"github.com/vitaliy-ukiru/fsm-telebot/storages/memory"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

var (
	InputSG        = fsm.NewStateGroup("book")
	InputTimeState = InputSG.New("time")
)

type Handler struct {
	log     *zap.SugaredLogger
	service *service.Service
	b       *tele.Bot
	cfg     *config.Config
}

func NewHandler(log *zap.SugaredLogger, service *service.Service, b *tele.Bot, cfg *config.Config) *Handler {
	return &Handler{log: log, service: service, b: b, cfg: cfg}
}

func (h *Handler) Register() {
	storage := memory.NewStorage()
	defer storage.Close()

	manager := fsm.NewManager(h.b, nil, storage, nil)

	h.b.Handle("/start", h.Hello)
	h.b.Handle(&keyboard.BtnVerify, h.VerifyUser)
	h.b.Handle(&keyboard.BtnBookRooms, h.SelectRoom)

	manager.Bind(&keyboard.BtnRoomOne, fsm.DefaultState, h.SelectTime)
	manager.Bind(&keyboard.BtnRoomTwo, fsm.DefaultState, h.SelectTime)
	manager.Bind(&keyboard.BtnTime, InputTimeState, h.BookRoom)
}
