package services

import (
	"errors"
	"go-example/internal/entities"
	"go-example/internal/log"

	"github.com/jinzhu/gorm"
)

//ProductService api controller of produces
type ProductService interface {
	FindAll(products *[]entities.Product, offset int, limit int, search string) error
	GetProduct(product *entities.Product, id string) error
	DeleteProduct(id string) error
}

type productService struct {
	db *gorm.DB
}

// NewProductService get product service instance
func NewProductService(db *gorm.DB) ProductService {
	return &productService{db}
}

func (p productService) FindAll(products *[]entities.Product, offset int, limit int, search string) error {
	// do stuff
	return p.db.Find(&products).Error
}

func (p productService) GetProduct(product *entities.Product, id string) error {
	// do stuff
	if err := p.db.Preload("Props").First(&product, entities.Product{Model: entities.Model{ID: id}}).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("product not found")
	}
	return nil
}

func (p productService) DeleteProduct(id string) error {
	log.Info("Delete product id=", id)
	tx := p.db.Begin()
	rs := tx.Delete(&entities.Product{Model: entities.Model{ID: id}}).Where("id=?", id)
	if rs.Error != nil {
		log.Error("fail to delete product:", rs.Error)
		tx.Rollback()
		return errors.New("can't delete product")
	}
	tx.Commit()
	return nil
}
