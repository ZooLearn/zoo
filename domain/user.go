package domain

import (
	"context"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
}
