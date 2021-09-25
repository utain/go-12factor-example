package v1

import (
	"go-example/internal/dto"
	"go-example/internal/entities"
	"go-example/internal/errors"
	"go-example/internal/log"
	"go-example/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// NewUserAPI create userService
func NewUserAPI(db *gorm.DB) UserAPI {
	return &userAPI{service: services.NewUserService(db)}
}

//UserAPI interface
type UserAPI interface {
	GetAllUser(c *gin.Context)
	GetUser(c *gin.Context)
}

// userAPI is a service private
type userAPI struct {
	service services.UserService
}

// GetAllUser return all User
func (p userAPI) GetAllUser(ctx *gin.Context) {
	users := new([]entities.User)
	if err := p.service.GetAllUser(users, 0, -1, ""); err != nil {
		ctx.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, dto.DataReply{Data: users})
}

// GetUser return only one User
func (p userAPI) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Debug("User id:", id)
	user := new(entities.User)
	if err := p.service.GetUser(user, id); err != nil {
		ctx.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, dto.DataReply{Data: user})
}
