package luma

import (
	"github.com/elizabeth-dev/plush_bot/internal/adapter"
	"github.com/elizabeth-dev/plush_bot/internal/types"
)

const TOKEN = ""
const BOT_ID = "luma"

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
	} else if req.Type == "command" {
		err = h.rawr(req.Message.Chat.Id, nil)
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

func (h *Handler) rawr(chatId int, replyTo *int) (err error) {
	payload := types.TelegramSendMessage{
		ChatId: chatId,
		Text:   RawrGenerator(),
	}

	if replyTo != nil {
		payload.ReplyTo = *replyTo
	}

	_, err = h.tg.SendMessage(TOKEN, payload)

	return err
}
