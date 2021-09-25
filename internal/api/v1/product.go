package v1

import (
	"go-example/internal/dto"
	"go-example/internal/entities"
	"go-example/internal/errors"
	"go-example/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
	var products []entities.Product
	p.service.FindAll(&products, 0, 10, "")
	ctx.JSON(http.StatusOK, dto.DataReply{Data: products})
}
func (p productAPI) GetProduct(ctx *gin.Context) {
	var product entities.Product
	if err := p.service.GetProduct(&product, ctx.Param("id")); err != nil {
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
