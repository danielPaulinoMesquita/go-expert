package database

import "github.com/devfullcycle/dan/goexpert/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(emaild string) (*entity.User, error)
}
