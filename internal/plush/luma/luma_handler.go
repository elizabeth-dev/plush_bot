package luma

import (
	"github.com/elizabeth-dev/plush_bot/internal/types"
)

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) Random() (*types.PlushResponse, error) {
	return &types.PlushResponse{Text: "Zzz"}, nil

}

func (h Handler) Mention(req types.PlushRequest) (*types.PlushResponse, error) {
	return &types.PlushResponse{Text: RawrGenerator()}, nil

}
