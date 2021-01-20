package users

type Service interface {
	GetUserSummary([]string) []*User
}

type Database interface {
	GetUserSummary([]string) []*User
}