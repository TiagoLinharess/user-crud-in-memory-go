package user

import "github.com/google/uuid"

type Id uuid.UUID

func (id Id) String() string {
	return uuid.UUID(id).String()
}

func ParseId(s string) (Id, error) {
	u, err := uuid.Parse(s)
	if err != nil {
		return Id{}, err
	}
	return Id(u), nil
}

type User struct {
	FirstName string
	LastName  string
	Biography string
}

type UserBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string
}

type UserResponse struct {
	Id        string
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string
}
