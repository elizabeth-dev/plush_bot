package types

type PlushHandler interface {
	Random() (*PlushResponse, error)
	Mention(req PlushRequest) (*PlushResponse, error)
}

type PlushRequest struct {
	BotId string
	Type  string // "random" OR "mention"
}

type PlushResponse struct {
	Text string
}
