package tigris

import (
	"github.com/elizabeth-dev/plush_bot/internal/adapter"
	"github.com/elizabeth-dev/plush_bot/internal/types"
)

const TOKEN = "5434967680:AAHvH0jmIa4t-is2GOqyzPnwn7XEHp44OPE"
const BOT_ID = "tigris"

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
