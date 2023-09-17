package tgbot

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
