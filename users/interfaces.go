package users

type Database interface {
	GetUserSummary([]string) []*User
}