package types

type TelegramUpdate struct {
	UpdateId int `json:"update_id"`
	Message  struct {
		MessageId int `json:"message_id"`
		From      struct {
			Id int `json:"id"`
		} `json:"from"`
		Chat struct {
			Id int `json:"id"`
		} `json:"chat"`
		Text string `json:"text"`
	} `json:"message"`
}

type TelegramSendMessage struct {
	Method string `json:"method"`
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}
