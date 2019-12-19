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

//ProductAPI api controller of produces
type ProductAPI interface {
	FindAll(*gin.Context)
	GetProduct(*gin.Context)
	DeleteProduct(*gin.Context)
}

type productAPI struct {
	productService services.ProductService
}

// NewProductAPI get product service instance
func NewProductAPI(db *gorm.DB, autoMigrate bool) ProductAPI {
	if autoMigrate {
		fmt.Println("Migrating Product model")
		db.AutoMigrate(&models.Product{}, &models.ProductProps{})
		db.Create(models.Product{
			Model: models.Model{ID: "1"}, Name: "MacBook Pro 13 2019", Code: "mbp132019", Price: 2999, Attr: models.AttrType{
				"with":   "23m",
				"height": "32m",
			},
			Props: []models.ProductProps{models.ProductProps{Model: models.Model{ID: "1"}, Key: "Name", Value: "MacDev"}},
		})
	}
	return &productAPI{productService: services.NewProductService(db)}
}

func (p productAPI) FindAll(ctx *gin.Context) {
	products := p.productService.FindAll(0, 10, "")
	ctx.JSON(http.StatusOK, products)
}
func (p productAPI) GetProduct(ctx *gin.Context) {
	product := p.productService.GetProduct(ctx.Param("id"))
	if reflect.DeepEqual(models.Product{}, *product) {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, product)
}
func (p productAPI) DeleteProduct(ctx *gin.Context) {
	err := p.productService.DeleteProduct(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusAccepted)
}
