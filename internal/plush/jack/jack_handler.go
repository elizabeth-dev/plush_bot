package jack

import (
	"github.com/elizabeth-dev/plush_bot/internal/adapter"
	"github.com/elizabeth-dev/plush_bot/internal/plush/sr_juan"
	"github.com/elizabeth-dev/plush_bot/internal/types"
	"time"
)

const TOKEN = "5818521061:AAEA7C-qugytp9cyUUhh2gxiok6sRDb1718"
const BOT_ID = "jack"

type Handler struct {
	tg *adapter.TelegramAdapter
}

func NewHandler(tg *adapter.TelegramAdapter) *Handler {
	return &Handler{tg: tg}
}

func (h *Handler) Handle(req *types.PlushRequest) (res *types.PlushRequest, err error) {
	if req.BotId != BOT_ID {
		return nil, nil
	}

	if req.Type == "command" {
		if req.Command == "oso_mono" {
			res, err := h.discuss(req.Message.Chat.Id)

			time.Sleep(time.Second)

			return &types.PlushRequest{
				Type: "internal", BotId: sr_juan.BOT_ID, Command: "discuss2", Message: types.PlushMessage{
					Id: *res,
					Chat: struct {
						Id    int
						Title string
					}{Id: req.Message.Chat.Id, Title: req.Message.Chat.Title},
				},
			}, err
		}
	}

	if req.Type == "internal" {

		if req.Command == "discuss3" {
			res, err := h.discuss3(req.Message.Chat.Id)

			time.Sleep(time.Second)

			return &types.PlushRequest{
				Type: "internal", BotId: sr_juan.BOT_ID, Command: "discuss4", Message: types.PlushMessage{
					Id: *res,
					Chat: struct {
						Id    int
						Title string
					}{Id: req.Message.Chat.Id, Title: req.Message.Chat.Title},
				},
			}, err
		}
		if req.Command == "discuss5" {
			res, err := h.discuss5(req.Message.Chat.Id)

			time.Sleep(time.Second)

			return &types.PlushRequest{
				Type: "internal", BotId: sr_juan.BOT_ID, Command: "discuss6", Message: types.PlushMessage{
					Id: *res,
					Chat: struct {
						Id    int
						Title string
					}{Id: req.Message.Chat.Id, Title: req.Message.Chat.Title},
				},
			}, err
		}
		if req.Command == "discuss7" {
			res, err := h.discuss7(req.Message.Chat.Id)

			time.Sleep(time.Second)

			return &types.PlushRequest{
				Type: "internal", BotId: sr_juan.BOT_ID, Command: "discuss8", Message: types.PlushMessage{
					Id: *res,
					Chat: struct {
						Id    int
						Title string
					}{Id: req.Message.Chat.Id, Title: req.Message.Chat.Title},
				},
			}, err
		}
		if req.Command == "discuss9" {
			err = h.discuss9(req.Message.Chat.Id)

			return nil, err
		}
	}

	return nil, err
}

func (h *Handler) discuss(chatId int) (msgId *int, err error) {
	payload := types.TelegramSendMessage{
		ChatId: chatId,
		Text:   "Soy el osito más cuqui",
	}

	msgId, err = h.tg.SendMessage(TOKEN, payload)

	return msgId, err
}

func (h *Handler) discuss3(chatId int) (msgId *int, err error) {
	payload := types.TelegramSendMessage{
		ChatId: chatId,
		Text:   "Mimimi, pues yo soy más grande y abrazable",
	}

	msgId, err = h.tg.SendMessage(TOKEN, payload)

	return msgId, err
}

func (h *Handler) discuss5(chatId int) (msgId *int, err error) {
	payload := types.TelegramSendMessage{
		ChatId: chatId,
		Text:   "Pues yo una vez viajé en metro, y saludé a la gente",
	}

	msgId, err = h.tg.SendMessage(TOKEN, payload)

	return msgId, err
}

func (h *Handler) discuss7(chatId int) (msgId *int, err error) {
	payload := types.TelegramSendMessage{
		ChatId: chatId,
		Text:   "Jopé... ¿Y si ambos somos súper cuquis?",
	}

	msgId, err = h.tg.SendMessage(TOKEN, payload)

	return msgId, err
}

func (h *Handler) discuss9(chatId int) (err error) {
	payload := types.TelegramSendMessage{
		ChatId: chatId,
		Text:   "Yupi 😊",
	}

	_, err = h.tg.SendMessage(TOKEN, payload)

	return err
}
