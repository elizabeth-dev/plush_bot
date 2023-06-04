package types

type TelegramUpdate struct {
	UpdateId     int              `json:"update_id"`
	Message      *TelegramMessage `json:"message"`
	MyChatMember *struct {
		Chat struct {
			Id    int    `json:"id"`
			Title string `json:"title"`
		} `json:"chat"`
		NewChatMember *struct {
			User   TelegramUser `json:"user"`
			Status string       `json:"status"`
		} `json:"new_chat_member"`
	} `json:"my_chat_member"`
}

type TelegramMessage struct {
	Id   int          `json:"message_id"`
	From TelegramUser `json:"from"`
	Chat struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
	} `json:"chat"`
	Text            string           `json:"text"`
	ReplyTo         *TelegramMessage `json:"reply_to_message"`
	NewParticipants []TelegramUser   `json:"new_chat_members"`
	Entities        []struct {
		Type   string `json:"type"`
		Offset int    `json:"offset"`
		Length int    `json:"length"`
	} `json:"entities"`
}

type TelegramUser struct {
	IsBot    bool   `json:"is_bot"`
	Name     string `json:"first_name"`
	Username string `json:"username"`
}

type TelegramSendMessage struct {
	ChatId      int                 `json:"chat_id"`
	Text        string              `json:"text"`
	ReplyTo     int                 `json:"reply_to_message_id,omitempty"`
	ReplyMarkup TelegramReplyMarkup `json:"reply_markup,omitempty"`
}

type TelegramReplyMarkup struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective"`
}

type TelegramSendVoice struct {
	ChatId int    `json:"chat_id"`
	Voice  string `json:"voice"`
}
