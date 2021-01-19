package database

type UserDto struct {
	ID    string
	Name  string
	Avatar  string
}

func NewUserDto(id, name, avatar string ) *UserDto {
	return &UserDto{
		ID:     id,
		Name:   name,
		Avatar: avatar,
	}
}