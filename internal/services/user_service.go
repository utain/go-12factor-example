package services

import (
	"errors"
	"go-example/internal/entities"
	"go-example/internal/log"

	"gorm.io/gorm"
)

// NewUserService create userService
func NewUserService(db *gorm.DB) UserService {
	return &userService{db}
}

//UserService interface
type UserService interface {
	GetAllUser(users *[]entities.User, offset int, limit int, search string) error
	GetUser(user *entities.User, id string) error
}

// userService is a service private
type userService struct {
	db *gorm.DB
}

// GetAllUser return all User
func (p userService) GetAllUser(users *[]entities.User, offset int, limit int, search string) error {
	if err := p.db.Find(&users).Error; err != nil {
		log.Error("GetAllUser", err)
		return nil
	}
	return nil
}

// GetUser return only one User
func (p userService) GetUser(user *entities.User, id string) error {
	if err := p.db.First(&user, &entities.User{Model: entities.Model{ID: id}}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		log.Error("GetUser", err)
		return errors.New("unknown error")
	}
	return nil
}
