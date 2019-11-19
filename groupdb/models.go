package groupdb

type Group struct {
	ID      int
	Name    string
	Members []Member `gorm:"many2many:group_members;"`
}

type Member struct {
	ID               int
	FirstName        string
	LastName         string
	Birthday         string
	TelegramUsername string
	Groups           []Group `gorm:"many2many:group_members;"`
}
