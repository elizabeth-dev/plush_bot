package horrifido

import (
	"github.com/elizabeth-dev/plush_bot/internal/adapter"
	"github.com/elizabeth-dev/plush_bot/internal/types"
)

const TOKEN = "5913311588:AAHMfWXAbu0jr7yx3Kf4bDDdk05UHFEax6Q"
const BOT_ID = "horrifido"

const SONG_ID = "AwACAgQAAxkDAAMCY8DPOhudQmLO4tSkYHwqEwpWv5YAAmgPAAI8KghSxxlY4j6upRMtBA"

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
	payload := types.TelegramSendVoice{
		ChatId: chatId,
		Voice:  SONG_ID,
	}

	err = h.tg.SendVoice(TOKEN, payload)

	return err
}
