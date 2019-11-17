package telegram

type Chat struct {
	Id int `json:"id"`
	Type string `json:"type"`
}

type Message struct {
	Id   int  `json:"message_id"`
	Timestamp int `json:"date"`
	Chat Chat `json:"chat"`
	Text string `json:"text"`
}

type Update struct {
	Id   int  `json:"update_id"`
	Message Message `json:"message"`
}

type SendMessageResponse struct {
	Method string `json:"method"`
	ChatId int `json:"chat_id"`
	Text string `json:"text"`
}