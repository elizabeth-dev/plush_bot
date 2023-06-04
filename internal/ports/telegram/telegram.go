package telegram

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/elizabeth-dev/plush_bot/internal/plush"
	"github.com/elizabeth-dev/plush_bot/internal/types"
)

type Handler struct {
	router plush.Router
}

func NewTelegramHandler(r plush.Router) *Handler {
	return &Handler{router: r}
}

func (h *Handler) ServeHTTP(_ http.ResponseWriter, req *http.Request) {
	bot := strings.Split(req.URL.EscapedPath(), "/")[1]

	if bot == "health" {
		return
	}

	var body = types.TelegramUpdate{}

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		// TODO: send crash alert
		log.Printf("ERROR: while unmarshalling request body: %v", err)

		return
	}

	fmt.Printf("BODY: %v\n", req.Body)

	if body.MyChatMember != nil && body.MyChatMember.NewChatMember.Status == "member" {
		h.router.Handle(
			&types.PlushRequest{
				BotId:   bot,
				Type:    "command",
				Command: "start",
				Message: types.PlushMessage{
					Chat: struct {
						Id    int
						Title string
					}{
						Id:    body.MyChatMember.Chat.Id,
						Title: body.MyChatMember.Chat.Title,
					},
				},
			},
		)
	}

	if body.MyChatMember != nil && body.MyChatMember.NewChatMember.Status == "left" {
		h.router.Handle(
			&types.PlushRequest{
				BotId:   bot,
				Type:    "command",
				Command: "stop",
				Message: types.PlushMessage{
					Chat: struct {
						Id    int
						Title string
					}{
						Id:    body.MyChatMember.Chat.Id,
						Title: body.MyChatMember.Chat.Title,
					},
				},
			},
		)
	}

	if body.Message != nil {
		request := types.PlushRequest{
			BotId: bot,
			Type:  "mention",
			Message: types.PlushMessage{
				Id: body.Message.Id,
				From: types.PlushUser{
					IsBot:    body.Message.From.IsBot,
					Name:     body.Message.From.Name,
					Username: body.Message.From.Username,
				},
				Chat: struct {
					Id    int
					Title string
				}{
					Id:    body.Message.Chat.Id,
					Title: body.Message.Chat.Title,
				},
				Text: body.Message.Text,
			},
		}

		if body.Message.ReplyTo != nil {
			request.Message.ReplyTo = &types.PlushMessage{
				Id: body.Message.ReplyTo.Id,
				From: types.PlushUser{
					IsBot:    body.Message.ReplyTo.From.IsBot,
					Name:     body.Message.ReplyTo.From.Name,
					Username: body.Message.ReplyTo.From.Username,
				},
				Chat: struct {
					Id    int
					Title string
				}{
					Id:    body.Message.ReplyTo.Chat.Id,
					Title: body.Message.ReplyTo.Chat.Title,
				},
				Text: body.Message.ReplyTo.Text,
			}
		}

		for _, participant := range body.Message.NewParticipants {
			request.Message.NewParticipants = append(
				request.Message.NewParticipants, types.PlushUser{
					IsBot:    participant.IsBot,
					Name:     participant.Name,
					Username: participant.Username,
				},
			)
		}

		for _, entity := range body.Message.Entities {
			if entity.Type == "bot_command" {
				request.Type = "command"
				request.Command = strings.Split(body.Message.Text[entity.Offset+1:entity.Offset+entity.Length], "@")[0]

			}
		}

		h.router.Handle(&request)
	}
}
