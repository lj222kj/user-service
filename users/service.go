package users

type service struct {
	database Database
}

func (s service) GetUserSummary(userIds []string) []*User {
	results := s.database.GetUserSummary(userIds)
	return results
}

func New(db Database) Service {
	return &service{
		database: db,
	}
}