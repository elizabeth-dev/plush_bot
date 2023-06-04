package main

import (
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/elizabeth-dev/plush_bot/internal/adapter"
	"github.com/elizabeth-dev/plush_bot/internal/plush"
	"github.com/elizabeth-dev/plush_bot/internal/ports/random"
	"github.com/elizabeth-dev/plush_bot/internal/ports/telegram"
)

func main() {
	telegramAdapter := adapter.NewTelegramAdapter(http.DefaultClient)
	mySession := session.Must(session.NewSession(
		aws.NewConfig().WithRegion(os.Getenv("DYNAMODB_REGION")).WithCredentials(credentials.NewEnvCredentials())))
	svc := dynamodb.New(mySession)

	router := plush.NewRouter(telegramAdapter, svc)
	telegramHandler := telegram.NewTelegramHandler(router)

	go random.NewRandomTimer(svc, &router)

	if err := http.ListenAndServe(":8090", telegramHandler); err != nil {
		panic(err)
	}
}
