package types

type PlushHandler interface {
	Handle(req *PlushRequest) (*PlushRequest, error)
}

type PlushRequest struct {
	BotId   string
	Type    string // "random", "mention", "internal" , or "command"
	Message PlushMessage
	Command string
}

type PlushMessage struct {
	Id   int
	From PlushUser
	Chat struct {
		Id    int
		Title string
	}
	Text            string
	ReplyTo         *PlushMessage
	NewParticipants []PlushUser
}

type PlushUser struct {
	IsBot    bool
	Name     string
	Username string
}
