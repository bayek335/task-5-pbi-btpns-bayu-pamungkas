package app

type UserLogging struct {
	Level        string
	UserID       string
	User         string
	Action       string
	BeforeUpdate string
	AfterUpdate  string
	Message      string
}
