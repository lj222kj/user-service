package web

import "user-service/users"

type UserService interface {
	GetUserSummary(userIds []string) []*users.User
}