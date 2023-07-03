package random

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/elizabeth-dev/plush_bot/internal/plush"
	"github.com/elizabeth-dev/plush_bot/internal/types"
)

var DYNAMODB_TABLE = os.Getenv("DYNAMODB_TABLE")
var ENV_MIN_SECONDS = os.Getenv("RANDOM_MIN_SECONDS")
var ENV_MAX_SECONDS = os.Getenv("RANDOM_MAX_SECONDS")

func NewRandomTimer(db *dynamodb.DynamoDB, router *plush.Router) {
	RANDOM_MIN_SECONDS, err := strconv.Atoi(ENV_MIN_SECONDS)

	if err != nil {
		RANDOM_MIN_SECONDS = 14400
	}

	RANDOM_MAX_SECONDS, err := strconv.Atoi(ENV_MAX_SECONDS)

	if err != nil {
		RANDOM_MAX_SECONDS = 21600
	}

	for {
		time.Sleep(time.Duration(rand.Intn(RANDOM_MAX_SECONDS-RANDOM_MIN_SECONDS)+RANDOM_MIN_SECONDS) * time.Second)

		fmt.Printf("RANDOM TIMER: %v\n", time.Now())

		chats, err := db.Scan(&dynamodb.ScanInput{
			TableName: &DYNAMODB_TABLE,
		})

		if err != nil {
			fmt.Printf("Error while scanning chats: %v\n", err)
			continue
		}

		var chatBots = make(map[int][]string)
		for _, chat := range chats.Items {
			chatId, err := strconv.Atoi(*chat["chatId"].N)
			if err != nil {
				fmt.Printf("Error while parsing %s to int: %v\n", *chat["chatId"].N, err)

				continue
			}

			chatBots[chatId] = []string{}

			if chat["bots"] != nil {
				for _, bot := range chat["bots"].SS {
					chatBots[chatId] = append(chatBots[chatId], *bot)
				}
			}
		}

		fmt.Printf("PARSED CHATBOTS: %v\n", chatBots)

		for chatId, bots := range chatBots {
			if len(bots) > 0 {
				router.Handle(
					&types.PlushRequest{
						BotId: bots[rand.Intn(len(bots))],
						Type:  "random",
						Message: types.PlushMessage{
							Chat: struct {
								Id    int
								Title string
							}{
								Id: chatId,
							},
						},
					},
				)
			}
		}
	}
}
