package application

import (
	"errors"
	"userCrud/models/user"

	"github.com/google/uuid"
)

type Application struct {
	data map[user.Id]user.User
}

func NewApplication() *Application {
	return &Application{
		data: make(map[user.Id]user.User),
	}
}

func (a *Application) FindAll() []user.UserResponse {
	users := make([]user.UserResponse, 0, len(a.data))
	for id, u := range a.data {
		u := user.UserResponse{
			Id:        id.String(),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Biography: u.Biography,
		}
		users = append(users, u)
	}
	return users
}

func (a *Application) FindById(id user.Id) (user.UserResponse, error) {
	u, ok := a.data[id]

	if !ok {
		return user.UserResponse{}, errors.New("user not found")
	}

	uJson := user.UserResponse{
		Id:        id.String(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Biography: u.Biography,
	}

	return uJson, nil
}

func (a *Application) Insert(u user.User) (user.User, user.Id) {
	id := user.Id(uuid.New())
	a.data[id] = u
	return u, id
}

func (a *Application) Update(u user.User, id user.Id) {
	a.data[id] = u
}

func (a *Application) Delete(id user.Id) {
	delete(a.data, id)
}
