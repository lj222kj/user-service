package database

import "user-service/users"

type Database interface {
	GetUserSummary([]string) []*users.User
	getUserSummary(string) (*users.User, error)
}