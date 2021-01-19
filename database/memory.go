package database

import (
	"fmt"
	"user-service/users"
)

type memory struct {
	users []*UserDto
}


func (m *memory) getUserSummary(id string) (*users.User, error) {
	fmt.Println("Test",m.users)
	for _, u := range m.users {
		if u.ID != id {
			continue
		}
		return &users.User{
			ID:     u.ID,
			Name:   u.Name,
			Avatar: u.Avatar,
		}, nil
	}
	return nil, fmt.Errorf("could not find user %s", id)
}

func (m *memory) GetUserSummary(userIds []string) []*users.User {
	var result []*users.User
	for _, id := range userIds {
		if user, err := m.getUserSummary(id); err == nil {
			result = append(result, user)
		}
	}
	return result
}

func New() Database {
	var users []*UserDto
	for i := 0; i < 10; i++{
		id := fmt.Sprint(i)
		users = append(users, &UserDto{
			ID:     id,
			Name:   "Jazmine_" + id,
		})
	}
	return &memory{users: users}
}