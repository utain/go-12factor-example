package v1

import (
	"fmt"
	"go-example/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// NewUserService create userService
func NewUserService(db *gorm.DB, autoMigrate bool) UserService {
	if autoMigrate {
		fmt.Println("Migrating User model")
		db.AutoMigrate(&models.User{})
		db.Create(models.User{
			Model: models.Model{ID: "1"}, Username: "utain", Email: "utain@gmail.com", Password: "P@55w0rd", Firstname: "Utain", Lastname: "Wongpreaw",
		})
	}
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

// GetAllUser return all User
func (p userService) GetAllUser(c *gin.Context) {
	var users []models.User
	p.db.Find(&users)
	c.JSON(http.StatusOK, users)
}

// GetUser return only one User
func (p userService) GetUser(c *gin.Context) {
	var user models.User
	p.db.First(&user, &models.User{Model: models.Model{ID: c.Param("id")}})
	if user == (models.User{}) {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
}
