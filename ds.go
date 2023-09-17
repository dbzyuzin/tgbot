package tgbot

type Callback struct {
	Data    string
	Message Message
}

type Message struct {
	ID     int
	User   User
	ChatID int64
	Text   string
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
