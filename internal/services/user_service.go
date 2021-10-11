package services

import (
	"errors"
	"go-example/internal/dto"
	"go-example/internal/entities"

	"gorm.io/gorm"
)

// NewUserService create userService
func NewUserService(db *gorm.DB) UserService {
	return &userService{db}
}

//UserService interface
type UserService interface {
	GetAllUser(page dto.Pageable) (*[]entities.User, error)
	GetUser(id string) (*entities.User, error)
}

// userService is a service private
type userService struct {
	db *gorm.DB
}

// GetAllUser return all User
func (p userService) GetAllUser(pageable dto.Pageable) (*[]entities.User, error) {
	users := new([]entities.User)
	p.db.Find(users)
	return users, nil
}

// GetUser return only one User
func (p userService) GetUser(id string) (*entities.User, error) {
	user := &entities.User{}
	if err := p.db.First(user, &entities.User{Model: entities.Model{ID: id}}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("unknown error")
	}
	return user, nil
}
