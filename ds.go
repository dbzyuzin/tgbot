package tgbot

type Message struct {
	ID     int
	User   User
	ChatID int64
}

type User struct {
	FirstName string
	LastName  string
	UserName  string
}