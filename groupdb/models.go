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

type Fundraiser struct {
	Group       Group
	GroupId     int
	Member      Member
	MemberId    int
	PhoneNumber string
	BanksList   string
}

type BirthdayCelebration struct {
	Group          Group
	GroupId        int
	Member         Member
	MemberId       int
	Fundraiser     Fundraiser
	FundraiserId   int
	TelegramChatId int
	Year           int
}
