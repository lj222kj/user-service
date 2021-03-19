package v2

import "user-service/proto"

type Database interface {
	GetUserSummaryV2([]string) []*proto.User
	GetAllUsers() []*proto.User
	GetAllUsersPaginated(lastEntry int64) []*proto.User
}
