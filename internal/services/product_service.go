package services

import (
	"go-example/internal/models"
	"log"

	"github.com/jinzhu/gorm"
)

//ProductService api controller of produces
type ProductService interface {
	FindAll(offset int, limit int, search string) *[]models.Product
	GetProduct(id string) *models.Product
	DeleteProduct(id string) error
}

type productService struct {
	db *gorm.DB
}

// NewProductService get product service instance
func NewProductService(db *gorm.DB) ProductService {
	return &productService{db}
}

func (p productService) FindAll(offset int, limit int, search string) *[]models.Product {
	var products []models.Product
	p.db.Find(&products)
	return &products
}

func (p productService) GetProduct(id string) *models.Product {
	var product models.Product
	p.db.Preload("Props").First(&product, models.Product{Model: models.Model{ID: id}})
	return &product
}

func (p productService) DeleteProduct(id string) error {
	log.Println("DeleteProduct with ID=" + id)
	tx := p.db.Begin()
	rs := tx.Delete(&models.Product{Model: models.Model{ID: id}}).Where("id=?", id)
	if rs.Error != nil {
		log.Fatal("[Error]:", rs.Error)
		tx.Rollback()
	}
	tx.Commit()
	return rs.Error
}
