package v1

import (
	"go-example/internal/models"
	"go-example/internal/services"
	"go-example/log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// NewUserAPI create userService
func NewUserAPI(db *gorm.DB) UserAPI {
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
	username := c.Param("name")
	log.Debug("GetUser:", username)
	user := p.userService.GetUser(username)
	if reflect.DeepEqual(*user, models.User{}) {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
}
