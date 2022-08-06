package telegram

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/elizabeth-dev/plush_bot/internal/plush"
	"github.com/elizabeth-dev/plush_bot/internal/types"
	"log"
	"strings"
)

var empty200 = events.LambdaFunctionURLResponse{StatusCode: 200}

type Handler struct {
	router plush.Router
}

func NewTelegramHandler(r plush.Router) Handler {
	return Handler{router: r}
}

func (h Handler) Handle(req events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	bot := strings.Split(req.RawPath, "/")[1]
	var body = types.TelegramUpdate{}

	if err := json.Unmarshal([]byte(req.Body), &body); err != nil {
		// TODO: send crash alert
		log.Printf("ERROR: while unmarshalling request body: %v", err)

		return empty200, nil
	}

	result, err := h.router.Handle(types.PlushRequest{
		BotId: bot,
		Type:  "mention",
	})

	if err != nil {
		// TODO: send crash alert
		log.Printf("ERROR: while plush handling request: %v", err)
		return empty200, nil
	}

	fmt.Printf("RESULT: %+v\n", result)

	res, err := json.MarshalIndent(types.TelegramSendMessage{
		Method: "sendMessage",
		ChatId: body.Message.Chat.Id,
		Text:   result.Text,
	}, "", "  ")

	if err != nil {
		// TODO: send crash alert
		log.Printf("ERROR: while marshalling plush response: %v", err)
		return empty200, nil
	}

	fmt.Printf("RESPONSE BODY: %s\n", res)

	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Body:       string(res),
	}, nil
}
