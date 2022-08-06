package plush

import (
	"fmt"
	"github.com/elizabeth-dev/plush_bot/internal/plush/luma"
	"github.com/elizabeth-dev/plush_bot/internal/types"
)

type Router struct {
	handlers map[string]types.PlushHandler
}

func NewRouter() Router {
	handlers := map[string]types.PlushHandler{"luma": luma.NewHandler()}
	return Router{handlers: handlers}
}

func (r Router) Handle(req types.PlushRequest) (*types.PlushResponse, error) {
	fmt.Printf("PLUSH: %s\n", req.BotId)
	fmt.Printf("REQUEST TYPE: %s\n", req.Type)

	switch req.Type {
	case "random":
		return r.handlers[req.BotId].Random()
	case "mention":
		return r.handlers[req.BotId].Mention(req)
	default:
		return nil, nil
	}
}
