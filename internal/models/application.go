package models

import (
	"errors"

	"github.com/google/uuid"
)

type Application struct {
	data map[uuid.UUID]User
}

func NewApplication() *Application {
	return &Application{
		data: make(map[uuid.UUID]User),
	}
}

func (a *Application) FindAll() map[uuid.UUID]User {
	return a.data
}

func (a *Application) FindById(id uuid.UUID) (User, error) {
	u, ok := a.data[id]

	if !ok {
		return User{}, errors.New("user not found")
	}

	return u, nil
}

func (a *Application) Insert(u User) (User, uuid.UUID) {
	id := uuid.New()
	a.data[id] = u
	return u, id
}

func (a *Application) Update(u User, id uuid.UUID) (User, uuid.UUID, error) {
	_, ok := a.data[id]

	if !ok {
		return User{}, id, errors.New("user not found")
	}

	a.data[id] = u

	return u, id, nil
}

func (a *Application) Delete(id uuid.UUID) (User, uuid.UUID, error) {
	u, ok := a.data[id]

	if !ok {
		return User{}, id, errors.New("user not found")
	}

	delete(a.data, id)

	return u, id, nil
}
