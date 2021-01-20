package users

type service struct {
	database Database
}

func (s service) GetUserSummary(userIds []string) []*User {
	return s.database.GetUserSummary(userIds)
}

func New(db Database) Service {
	return &service{
		database: db,
	}
}