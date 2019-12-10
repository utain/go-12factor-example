package v1

import (
	"fmt"
	"go-example/app/db/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetService for get userService
func GetService(db *gorm.DB) UserService {
	return &userService{db}
}

//UserService interface
type UserService interface {
	GetAllUser(c *gin.Context)
	GetUser(c *gin.Context)
}

// userService is a service private
type userService struct {
	db *gorm.DB
}

func (p userService) SetDB(db *gorm.DB) {
	p.db = db
}

func (p userService) getDB() *gorm.DB {
	return p.db
}

// Migrate auto migrate db schema
func Migrate(db *gorm.DB) {
	fmt.Println("Migrating User model")
	db.AutoMigrate(&model.User{})
}

// GetAllUser return all User
func (p userService) GetAllUser(c *gin.Context) {
	var users []model.User
	p.getDB().Find(&users)
	c.JSON(200, users)
}

// GetUser return only one User
func (p userService) GetUser(c *gin.Context) {
	var user model.User
	p.getDB().First(&user, &model.User{Username: c.Param("id")})
	if user != (model.User{}) {
		c.JSON(200, user)
	}
	c.Status(404)
}
