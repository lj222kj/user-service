package database

import (
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"user-service/proto"
	"user-service/users"
)

type memory struct {
	users []*UserDto
}

func (m *memory) GetAllUsersPaginated(lastEntry int64) []*proto.User {
	var result []*proto.User
	var userLen int64
	userLen = int64(len(m.users))
	if userLen == lastEntry || lastEntry > userLen{
		return result
	}
	for _, user := range m.users[lastEntry:lastEntry + 19] {
		result = append(result, &proto.User{
			Id:        user.ID,
			Name:      user.Name,
			Avatar:    user.Avatar,
			CreatedAt: timestamppb.Now(),
		})
	}
	return result
}

func (m *memory) GetUserSummaryV2(ids []string) []*proto.User {
	var result []*proto.User
	for _, id := range ids {
		if user, err := m.getUserSummary(id); err == nil {
			result = append(result, &proto.User{
				Id:     user.ID,
				Name:   user.Name,
				Avatar: user.Avatar,
				CreatedAt: timestamppb.Now(),
			})
		}
	}
	return result
}


func (m *memory) GetAllUsers() []*proto.User {
	var result []*proto.User
	for _, user := range m.users {
		result = append(result, &proto.User{
			Id:        user.ID,
			Name:      user.Name,
			Avatar:    user.Avatar,
			CreatedAt: timestamppb.Now(),
		})
	}
	return result
}

func (m *memory) getUserSummary(id string) (*users.User, error) {
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

func New() *memory {
	var users []*UserDto
	for i := 0; i < 50; i++{
		id := fmt.Sprint(i)
		users = append(users, NewUserDto(id, "Jazmine_" + id, ""))
	}
	return &memory{users: users}
}