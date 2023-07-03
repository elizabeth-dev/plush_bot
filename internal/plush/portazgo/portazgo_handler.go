package portazgo

import (
	"strings"

	"github.com/elizabeth-dev/plush_bot/internal/adapter"
	"github.com/elizabeth-dev/plush_bot/internal/types"
)

const TOKEN = ""
const BOT_ID = "portazgo"

const PASSWORD = "miau"
const INCORRECT_PASS_STR = "¡Incorrecto! Vuelve a intentarlo..."

type Handler struct {
	tg *adapter.TelegramAdapter
}

func NewHandler(tg *adapter.TelegramAdapter) *Handler {
	return &Handler{tg: tg}
}

func (h *Handler) Handle(req *types.PlushRequest) (_ *types.PlushRequest, err error) {
	if req.BotId != BOT_ID {
		return nil, err
	}

	if req.Type == "random" {
		err = h.sleep(req.Message.Chat.Id)
	} else {
		if len(req.Message.NewParticipants) > 0 {
			for _, user := range req.Message.NewParticipants {
				if !user.IsBot {
					err = h.askPassword(req.Message.Chat.Id, user.Username)
				}
			}

			return nil, err
		}

		if req.Message.ReplyTo != nil && req.Message.ReplyTo.From.IsBot && (strings.Contains(
			req.Message.ReplyTo.Text,
			"Alto ahí",
		) || req.Message.ReplyTo.Text == INCORRECT_PASS_STR) {
			err = h.verifyPassword(req.Message.Chat.Id, req.Message.Id, req.Message.Text, req.Message.From.Name)
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

func (h *Handler) askPassword(chatId int, username string) (err error) {
	payload := types.TelegramSendMessage{
		ChatId: chatId,
		Text:   "¡Alto ahí @" + username + "! ¿Contraseña?",
		ReplyMarkup: types.TelegramReplyMarkup{
			ForceReply: true,
			Selective:  true,
		},
	}

	_, err = h.tg.SendMessage(TOKEN, payload)

	return err
}

func (h *Handler) verifyPassword(chatId int, messageId int, password string, userName string) (err error) {
	var payload types.TelegramSendMessage

	if strings.Contains(strings.ToLower(password), PASSWORD) {
		payload = types.TelegramSendMessage{
			ChatId:  chatId,
			Text:    "¡Correcto! Adelante, " + userName,
			ReplyTo: messageId,
		}
	} else {
		payload = types.TelegramSendMessage{
			ChatId:  chatId,
			Text:    INCORRECT_PASS_STR,
			ReplyTo: messageId,
			ReplyMarkup: types.TelegramReplyMarkup{
				ForceReply: true,
				Selective:  true,
			},
		}
	}

	_, err = h.tg.SendMessage(TOKEN, payload)

	return err
}
