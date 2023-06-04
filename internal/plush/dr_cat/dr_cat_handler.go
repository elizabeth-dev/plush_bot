package dr_cat

import (
	"github.com/elizabeth-dev/plush_bot/internal/adapter"
	"github.com/elizabeth-dev/plush_bot/internal/types"
	"math/rand"
)

const TOKEN = "5825383094:AAEsJbI6H7W-2-fVHQ0Uk-osy9LoZpDS9wE"
const BOT_ID = "dr_cat"

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
		num := rand.Intn(2)
		if num == 0 {
			err = h.sleep(req.Message.Chat.Id)
		} else {
			err = h.terminal(req.Message.Chat.Id)
		}
	}

	if req.Type == "command" {
		if req.Command == "consejo" {
			err = h.advice(req.Message.Chat.Id)
		}
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

func (h *Handler) terminal(chatId int) (err error) {
	_, err = h.tg.SendMessage(
		TOKEN, types.TelegramSendMessage{
			ChatId: chatId,
			Text:   "It's terminal",
		},
	)

	return err
}

func (h *Handler) advice(chatId int) (err error) {
	payload := types.TelegramSendMessage{
		ChatId: chatId,
		Text:   MEDICAL_ADVICE[rand.Intn(len(MEDICAL_ADVICE))],
	}

	_, err = h.tg.SendMessage(TOKEN, payload)

	return err
}
