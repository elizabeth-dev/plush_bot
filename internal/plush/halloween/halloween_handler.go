package halloween

import (
	"github.com/elizabeth-dev/plush_bot/internal/adapter"
	"github.com/elizabeth-dev/plush_bot/internal/types"
)

const TOKEN = ""
const BOT_ID = "halloween"

const SONG_ID = "AwACAgQAAxkDAAMCY8EsmrvGsd2o5x6LDMGlkiQ3VNEAAp0NAAI8KhBSQq8S8pDfycktBA"

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

	if req.Type == "random" {
		err = h.sleep(req.Message.Chat.Id)
	} else if req.Command == "canta" {
		err = h.sing(req.Message.Chat.Id)
	}

	return nil, err
}
func (h *Handler) sleep(chatId int) (err error) {
	_, err = h.tg.SendMessage(
		TOKEN, types.TelegramSendMessage{
			ChatId: chatId,
			Text:   "zzZZzzZZzzZZ",
		},
	)

	return err
}

func (h *Handler) sing(chatId int) (err error) {
	payload := types.TelegramSendVoice{
		ChatId: chatId,
		Voice:  SONG_ID,
	}

	err = h.tg.SendVoice(TOKEN, payload)

	return err
}
