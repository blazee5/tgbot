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
	InputRoomGroup = fsm.NewStateGroup("book")
	InputTimeState = InputRoomGroup.New("time")

	InputRegGroup  = fsm.NewStateGroup("reg")
	InputFirstName = InputRegGroup.New("firstname")
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

	manager.Bind("/start", fsm.DefaultState, h.Start)
	manager.Bind(tele.OnText, InputFirstName, h.RegisterUser)

	h.b.Handle(&keyboard.BtnVerify, h.VerifyUser)
	h.b.Handle(&keyboard.BtnBookRooms, h.SelectRoom)
	h.b.Handle(&keyboard.BtnGetBooks, h.GetBooks)
	h.b.Handle(&keyboard.BtnCancelBook, h.CancelBook)

	manager.Bind(&keyboard.BtnRoomOne, fsm.DefaultState, h.SelectTime)
	manager.Bind(&keyboard.BtnRoomTwo, fsm.DefaultState, h.SelectTime)
	manager.Bind(&keyboard.BtnTime, InputTimeState, h.BookRoom)
}
