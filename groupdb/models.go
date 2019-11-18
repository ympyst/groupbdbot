package groupdb

type Group struct {
	Id int
	Name string
}

type Member struct {
	Id               int
	GroupId          int
	FirstName        string
	LastName         string
	Birthday         string
	TelegramUsername string
}
