package v1

import (
	"go-example/internal/dto"
	"go-example/internal/errors"
	"go-example/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

// GetAllUser godoc
// @Title GetAllUser
// @Summary Get all users information
// @Description List users in server
// @ID get-all-users
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.DataReply{data=[]entities.User}
// @Failure 500 {object} dto.ErrorReply "Unknown error"
// @Router /users [get]
func (p *userAPI) GetAllUser(ctx *gin.Context) {
	users, err := p.service.GetAllUser(dto.Pageable{})
	if err != nil {
		ctx.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, dto.DataReply{Data: users})
}

// GetUser return only one User
func (p *userAPI) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := p.service.GetUser(id)
	if err != nil {
		ctx.Error(errors.NewError(http.StatusNotFound, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, dto.DataReply{Data: user})
}
