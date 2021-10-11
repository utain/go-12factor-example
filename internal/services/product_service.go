package services

import (
	"errors"
	"go-example/internal/dto"
	"go-example/internal/entities"
	"go-example/internal/log"

	"gorm.io/gorm"
)

//ProductService api controller of produces
type ProductService interface {
	FindAll(pageable dto.Pageable) (*[]entities.Product, error)
	GetProduct(id string) (*entities.Product, error)
	DeleteProduct(id string) error
}

type productService struct {
	db *gorm.DB
}

// NewProductService get product service instance
func NewProductService(db *gorm.DB) ProductService {
	return &productService{db}
}

func (p productService) FindAll(pageable dto.Pageable) (*[]entities.Product, error) {
	// do stuff
	products := new([]entities.Product)
	p.db.Find(products)
	return products, nil
}

func (p productService) GetProduct(id string) (*entities.Product, error) {
	// do stuff
	product := &entities.Product{}
	if err := p.db.Preload("Props").First(&product, entities.Product{Model: entities.Model{ID: id}}).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("product not found")
	}
	return product, nil
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
