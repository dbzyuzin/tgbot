package tgbot

type Callback struct {
	Data    string
	Message Message
}

type Message struct {
	ID       int
	User     User
	ChatID   int64
	Text     string
	HTMLText string
	Entities []MessageEntity

	IsTextOnly bool
}

type MessageEntity struct {
	Type   string // "bold", "italic", "code", "pre", "text_link", etc.
	Offset int    // Начальная позиция
	Length int    // Длина
	URL    string // Для text_link
}

type User struct {
	ID        int64
	FirstName string
	LastName  string
	UserName  string
}

type Button struct {
	Text string
	Data string
}
