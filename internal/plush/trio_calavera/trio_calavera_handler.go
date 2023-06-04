package trio_calavera

import (
	"github.com/elizabeth-dev/plush_bot/internal/adapter"
	"github.com/elizabeth-dev/plush_bot/internal/types"
	"math/rand"
)

const TOKEN = "5948803373:AAF7IrMFCjmzi-Pkodxbgj29bflTSv0cpb0"
const BOT_ID = "trio_calavera"

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
		if req.Command == "curiosidad" {
			err = h.trivia(req.Message.Chat.Id)
		}
	}

	return nil, err
}

func (h *Handler) trivia(chatId int) (err error) {
	payload := types.TelegramSendMessage{
		ChatId: chatId,
		Text:   RANDOM_TRIVIA[rand.Intn(len(RANDOM_TRIVIA))],
	}

	_, err = h.tg.SendMessage(TOKEN, payload)

	return err
}
