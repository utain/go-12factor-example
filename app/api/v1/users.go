package v1

import (
	"fmt"
	"go-example/app/models"
	"go-example/app/services"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// NewUserAPI create userService
func NewUserAPI(db *gorm.DB, autoMigrate bool) UserAPI {
	if autoMigrate {
		fmt.Println("Migrating User model")
		db.AutoMigrate(&models.User{})
		db.Create(models.User{
			Model: models.Model{ID: "1"}, Username: "utain", Email: "utain@gmail.com", Password: "P@55w0rd", Firstname: "Utain", Lastname: "Wongpreaw",
		})
	}
	return &userAPI{userService: services.NewUserService(db)}
}

//UserAPI interface
type UserAPI interface {
	GetAllUser(c *gin.Context)
	GetUser(c *gin.Context)
}

// userAPI is a service private
type userAPI struct {
	userService services.UserService
}

// GetAllUser return all User
func (p userAPI) GetAllUser(c *gin.Context) {
	users := p.userService.GetAllUser(0, -1, "")
	c.JSON(http.StatusOK, users)
}

// GetUser return only one User
func (p userAPI) GetUser(c *gin.Context) {
	user := p.userService.GetUser(c.Param("name"))
	if reflect.DeepEqual(*user, models.User{}) {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
}
