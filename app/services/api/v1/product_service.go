package v1

import (
	"fmt"
	"go-example/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//ProductService service of produces
type ProductService interface {
	FindAll(*gin.Context)
	GetProduct(*gin.Context)
	DeleteProduct(*gin.Context)
}
type productServ struct {
	db *gorm.DB
}

// NewProductService get product service instance
func NewProductService(db *gorm.DB, autoMigrate bool) ProductService {
	if autoMigrate {
		fmt.Println("Migrating Product model")
		db.AutoMigrate(&models.Product{})
		db.Create(models.Product{
			Model: models.Model{ID: "1"}, Name: "MacBook Pro 13 2019", Code: "mbp132019", Price: 2999,
		})
	}
	return &productServ{db}
}

func (p productServ) FindAll(ctx *gin.Context) {
	var products []models.Product
	p.db.Find(&products)
	ctx.JSON(http.StatusOK, products)
}
func (p productServ) GetProduct(ctx *gin.Context) {
	var product models.Product
	p.db.First(&product, models.Product{Model: models.Model{ID: ctx.Param("id")}})
	if (models.Product{}) == product {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, product)
}
func (p productServ) DeleteProduct(ctx *gin.Context) {
	p.db.Delete(models.Product{}, models.Product{Model: models.Model{ID: ctx.Param("id")}})
	ctx.Status(http.StatusAccepted)
}
