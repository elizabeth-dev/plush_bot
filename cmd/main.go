package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/elizabeth-dev/plush_bot/internal/plush"
	"github.com/elizabeth-dev/plush_bot/internal/ports/telegram"
	"strings"
)

func main() {
	router := plush.NewRouter()
	telegramHandler := telegram.NewTelegramHandler(router)

	lambda.Start(func(req events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
		eventJson, _ := json.MarshalIndent(req, "", "  ")
		fmt.Printf("EVENT: %s\n", eventJson)

		handler := strings.Split(req.RawPath, "/")[2]
		fmt.Printf("HANDLER: %s\n", handler)

		switch handler {
		case "telegram":
			return telegramHandler.Handle(req)
		//case "cron":
		default:
			return events.LambdaFunctionURLResponse{StatusCode: 200}, nil
		}
	})
}
