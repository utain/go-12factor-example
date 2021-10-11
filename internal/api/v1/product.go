package v1

import (
	"go-example/internal/dto"
	"go-example/internal/errors"
	"go-example/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//ProductAPI api controller of produces
type ProductAPI interface {
	FindAll(*gin.Context)
	GetProduct(*gin.Context)
	DeleteProduct(*gin.Context)
}

type productAPI struct {
	service services.ProductService
}

// NewProductAPI get product service instance
func NewProductAPI(db *gorm.DB) ProductAPI {
	return &productAPI{service: services.NewProductService(db)}
}

func (p productAPI) FindAll(ctx *gin.Context) {
	products, _ := p.service.FindAll(dto.Pageable{})
	ctx.JSON(http.StatusOK, dto.DataReply{Data: products})
}

func (p productAPI) GetProduct(ctx *gin.Context) {
	product, err := p.service.GetProduct(ctx.Param("id"))
	if err != nil {
		ctx.Error(errors.NewError(http.StatusNotFound, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, dto.DataReply{Data: product})
}

func (p productAPI) DeleteProduct(ctx *gin.Context) {
	if err := p.service.DeleteProduct(ctx.Param("id")); err != nil {
		ctx.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.Status(http.StatusAccepted)
}
