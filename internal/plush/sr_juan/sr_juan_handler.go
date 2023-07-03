package sr_juan

import (
	"fmt"
	"strings"
	"time"

	"github.com/elizabeth-dev/plush_bot/internal/adapter"
	"github.com/elizabeth-dev/plush_bot/internal/types"
)

const TOKEN = ""
const BOT_ID = "sr_juan"

type Handler struct {
	tg *adapter.TelegramAdapter
}

func NewHandler(tg *adapter.TelegramAdapter) *Handler {
	return &Handler{tg: tg}
}

func (h *Handler) Handle(req *types.PlushRequest) (_ *types.PlushRequest, err error) {
	fmt.Printf("Handling request: %+v\n", req)
	if req.BotId != BOT_ID {
		return nil, nil
	}

	if req.Type == "command" {
		if req.Command == "edad" {
			err = h.sayAge(req.Message.Chat.Id)
			return nil, err
		}
	}

	if req.Type == "internal" {
		if req.Command == "discuss2" {
			res, err := h.discuss2(req.Message.Chat.Id)

			time.Sleep(time.Second)

			return &types.PlushRequest{
				Type: "internal", BotId: "jack", Command: "discuss3", Message: types.PlushMessage{
					Id: *res,
					Chat: struct {
						Id    int
						Title string
					}{Id: req.Message.Chat.Id, Title: req.Message.Chat.Title},
				},
			}, err
		}
		if req.Command == "discuss4" {
			res, err := h.discuss4(req.Message.Chat.Id)

			time.Sleep(time.Second)

			return &types.PlushRequest{
				Type: "internal", BotId: "jack", Command: "discuss5", Message: types.PlushMessage{
					Id: *res,
					Chat: struct {
						Id    int
						Title string
					}{Id: req.Message.Chat.Id, Title: req.Message.Chat.Title},
				},
			}, err
		}
		if req.Command == "discuss6" {
			res, err := h.discuss6(req.Message.Chat.Id)

			time.Sleep(time.Second)

			return &types.PlushRequest{
				Type: "internal", BotId: "jack", Command: "discuss7", Message: types.PlushMessage{
					Id: *res,
					Chat: struct {
						Id    int
						Title string
					}{Id: req.Message.Chat.Id, Title: req.Message.Chat.Title},
				},
			}, err
		}
		if req.Command == "discuss8" {
			res, err := h.discuss8(req.Message.Chat.Id)

			time.Sleep(time.Second)

			return &types.PlushRequest{
				Type: "internal", BotId: "jack", Command: "discuss9", Message: types.PlushMessage{
					Id: *res,
					Chat: struct {
						Id    int
						Title string
					}{Id: req.Message.Chat.Id, Title: req.Message.Chat.Title},
				},
			}, err
		}
	}

	return nil, err
}

func (h *Handler) sayAge(chatId int) (err error) {
	payload := types.TelegramSendMessage{
		ChatId: chatId,
		Text:   AgeGenerator(),
	}

	_, err = h.tg.SendMessage(TOKEN, payload)

	return err
}

func (h *Handler) discuss2(chatId int) (msgId *int, err error) {
	payload := types.TelegramSendMessage{
		ChatId: chatId,
		Text:   "No, ese soy yo >:( Además, yo soy más peque",
	}

	msgId, err = h.tg.SendMessage(TOKEN, payload)

	return msgId, err
}

func (h *Handler) discuss4(chatId int) (msgId *int, err error) {
	payload := types.TelegramSendMessage{
		ChatId: chatId,
		Text:   "Pues yo llevo más tiempo con Dafne",
	}

	msgId, err = h.tg.SendMessage(TOKEN, payload)

	return msgId, err
}

func (h *Handler) discuss6(chatId int) (msgId *int, err error) {
	payload := types.TelegramSendMessage{
		ChatId: chatId,
		Text:   "Pues yo soy más cuqui porque " + strings.ToLower(AgeGenerator()),
	}

	msgId, err = h.tg.SendMessage(TOKEN, payload)

	return msgId, err
}

func (h *Handler) discuss8(chatId int) (msgId *int, err error) {
	payload := types.TelegramSendMessage{
		ChatId: chatId,
		Text:   "Hmmm... De acuerdo, me parece bien",
	}

	msgId, err = h.tg.SendMessage(TOKEN, payload)

	return msgId, err
}
