package liz

import (
	"github.com/elizabeth-dev/plush_bot/internal/adapter"
	"github.com/elizabeth-dev/plush_bot/internal/types"
	"math/rand"
)

const TOKEN = "5666394080:AAFgSXHhVPuVJCJY7cdGscCu6nWl7mNXH_M"
const BOT_ID = "liz"

type Handler struct {
	tg *adapter.TelegramAdapter
}

func NewHandler(tg *adapter.TelegramAdapter) *Handler {
	return &Handler{tg: tg}
}

func (h *Handler) Handle(req *types.PlushRequest) (_ *types.PlushRequest, err error) {
	if req.BotId != BOT_ID {
		return nil, nil
	}

	if req.Type == "command" {
		if req.Command == "canta" {
			err = h.sing(req.Message.Chat.Id)
		}
	}

	return nil, err
}

func (h *Handler) sing(chatId int) (err error) {
	payload := types.TelegramSendMessage{
		ChatId: chatId,
		Text:   SONG_FRAGMENTS[rand.Intn(len(SONG_FRAGMENTS))],
	}

	_, err = h.tg.SendMessage(TOKEN, payload)

	return err
}
