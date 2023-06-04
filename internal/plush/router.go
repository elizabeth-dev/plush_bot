package plush

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/elizabeth-dev/plush_bot/internal/adapter"
	"github.com/elizabeth-dev/plush_bot/internal/plush/dr_cat"
	"github.com/elizabeth-dev/plush_bot/internal/plush/halloween"
	"github.com/elizabeth-dev/plush_bot/internal/plush/horrifido"
	"github.com/elizabeth-dev/plush_bot/internal/plush/jack"
	"github.com/elizabeth-dev/plush_bot/internal/plush/liz"
	"github.com/elizabeth-dev/plush_bot/internal/plush/luma"
	"github.com/elizabeth-dev/plush_bot/internal/plush/portazgo"
	"github.com/elizabeth-dev/plush_bot/internal/plush/sr_juan"
	"github.com/elizabeth-dev/plush_bot/internal/plush/tigris"
	"github.com/elizabeth-dev/plush_bot/internal/plush/trio_calavera"
	"github.com/elizabeth-dev/plush_bot/internal/types"
)

var DYNAMODB_TABLE = os.Getenv("DYNAMODB_TABLE")

type Router struct {
	handlers map[string]types.PlushHandler
	db       *dynamodb.DynamoDB
}

func NewRouter(tg *adapter.TelegramAdapter, db *dynamodb.DynamoDB) Router {
	handlers := map[string]types.PlushHandler{
		luma.BOT_ID:          luma.NewHandler(tg),
		portazgo.BOT_ID:      portazgo.NewHandler(tg),
		sr_juan.BOT_ID:       sr_juan.NewHandler(tg),
		trio_calavera.BOT_ID: trio_calavera.NewHandler(tg),
		dr_cat.BOT_ID:        dr_cat.NewHandler(tg),
		liz.BOT_ID:           liz.NewHandler(tg),
		horrifido.BOT_ID:     horrifido.NewHandler(tg),
		halloween.BOT_ID:     halloween.NewHandler(tg),
		tigris.BOT_ID:        tigris.NewHandler(tg),
		jack.BOT_ID:          jack.NewHandler(tg),
	}
	return Router{handlers: handlers, db: db}
}

func (r *Router) Handle(req *types.PlushRequest) {
	fmt.Printf("PLUSH: %s\n", req.BotId)
	fmt.Printf("REQUEST: %v\n", req)

	if req.Type == "command" && req.Command == "start" {
		r.saveBot(req)
	}

	if req.Type == "command" && req.Command == "stop" {
		r.removeBot(req)
	}

	for plush, handler := range r.handlers {
		res, err := handler.Handle(req)

		if err != nil {
			fmt.Printf("ERROR IN BOT %s: %v\n", plush, err)
		}

		if res != nil {
			r.Handle(res)
		}
	}
}

func (r *Router) saveBot(req *types.PlushRequest) {
	chatIdKey := map[string]*dynamodb.AttributeValue{
		"chatId": {
			N: aws.String(strconv.Itoa(req.Message.Chat.Id)),
		},
	}

	result, err := r.db.GetItem(&dynamodb.GetItemInput{
		TableName: &DYNAMODB_TABLE,
		Key:       chatIdKey,
	})

	if err != nil {
		fmt.Printf("ERROR: While saving bot %s querying chatId %d: %v\n", req.BotId, req.Message.Chat.Id, err)
		return
	}

	if result.Item != nil {
		if result.Item["bots"] != nil {
			bots := result.Item["bots"].SS
			for _, bot := range bots {
				if *bot == req.BotId {
					return
				}
			}
		}

		_, err = r.db.UpdateItem(&dynamodb.UpdateItemInput{
			TableName:        &DYNAMODB_TABLE,
			Key:              chatIdKey,
			UpdateExpression: aws.String("ADD bots :b"),
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":b": {
					SS: []*string{
						aws.String(req.BotId),
					},
				},
			},
		})

		if err != nil {
			fmt.Printf("ERROR: While saving bot %s appending to chatId %d: %v\n", req.BotId, req.Message.Chat.Id, err)
		}
		return
	}

	_, err = r.db.PutItem(&dynamodb.PutItemInput{
		TableName: &DYNAMODB_TABLE,
		Item: map[string]*dynamodb.AttributeValue{
			"chatId": {
				N: aws.String(strconv.Itoa(req.Message.Chat.Id)),
			},
			"bots": {
				SS: []*string{aws.String(req.BotId)},
			},
		},
	})
	if err != nil {
		fmt.Printf("ERROR: While saving bot %s creating chatId %d: %v\n", req.BotId, req.Message.Chat.Id, err)
	}
}

func (r *Router) removeBot(req *types.PlushRequest) {
	_, err := r.db.UpdateItem(&dynamodb.UpdateItemInput{
		TableName:        &DYNAMODB_TABLE,
		Key:              map[string]*dynamodb.AttributeValue{"chatId": {N: aws.String(strconv.Itoa(req.Message.Chat.Id))}},
		UpdateExpression: aws.String("DELETE bots :b"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":b": {
				SS: []*string{
					aws.String(req.BotId),
				},
			},
		},
	})

	if err != nil {
		fmt.Printf("ERROR REMOVING BOT: %v\n", err)
	}
}
