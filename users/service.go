package users

type service struct {
	database Database
}

func (s service) GetUserSummary(userIds []string) []*User {
	results := s.database.GetUserSummary(userIds)
	return results
}

func New(db Database) service {
	return service {
		database: db,
	}
}