package application

import (
	"errors"
	"userCrud/models/user"

	"github.com/google/uuid"
)

type Application struct {
	data map[uuid.UUID]user.User
}

func NewApplication() *Application {
	return &Application{
		data: make(map[uuid.UUID]user.User),
	}
}

func (a *Application) FindAll() map[uuid.UUID]user.User {
	return a.data
}

func (a *Application) FindById(id uuid.UUID) (user.User, error) {
	u, ok := a.data[id]

	if !ok {
		return user.User{}, errors.New("user not found")
	}

	return u, nil
}

func (a *Application) Insert(u user.User) (user.User, uuid.UUID) {
	id := uuid.New()
	a.data[id] = u
	return u, id
}

func (a *Application) Update(u user.User, id uuid.UUID) (user.User, uuid.UUID, error) {
	_, ok := a.data[id]

	if !ok {
		return user.User{}, id, errors.New("user not found")
	}

	a.data[id] = u

	return u, id, nil
}

func (a *Application) Delete(id uuid.UUID) (user.User, uuid.UUID, error) {
	u, ok := a.data[id]

	if !ok {
		return user.User{}, id, errors.New("user not found")
	}

	delete(a.data, id)

	return u, id, nil
}
