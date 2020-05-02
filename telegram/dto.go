package telegram

type User struct {
	Id int `json:"id"`
	IsBot bool `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
}

type Chat struct {
	Id int `json:"id"`
	Type string `json:"type"`
	Username string `json:"username"`
}

type Message struct {
	Id   int  `json:"message_id"`
	Timestamp int `json:"date"`
	UserFrom User `json:"from"`
	Chat Chat `json:"chat"`
	Text string `json:"text"`
	Entities []MessageEntity `json:"entities"`
}

type MessageEntity struct {
	Type string `json:"type"`
	Offset int `json:"offset"`
	Length int `json:"length"`
	Url string `json:"url"`
	User User `json:"user"`
}

type Update struct {
	Id   int  `json:"update_id"`
	Message *Message `json:"message,omitempty"`
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
}

type CallbackQuery struct {
	Id       string `json:"id"`
	UserFrom User   `json:"from"`
	Data     string `json:"data"`
} 

type SendMessageResponse struct {
	Method string `json:"method"`
	ChatId int `json:"chat_id"`
	Text string `json:"text"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type AnswerCallbackQueryResponse struct {
	CallbackQueryId string `json:"callback_query_id"`
	Text string `json:"text"`
}

type ReplyKeyboardMarkup struct {
	Keyboard [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard bool `json:"resize_keyboard"`
	OneTimeKeyboard bool `json:"one_time_keyboard"`
	Selective bool `json:"selective"`
}

type KeyboardButton struct {
	Text string `json:"text"`
}

type InlineKeyboardMarkup struct {
	Keyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text string `json:"text"`
	CallbackData string `json:"callback_data"`
}