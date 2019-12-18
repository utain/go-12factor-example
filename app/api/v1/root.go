package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// RegisterRouterAPIV1 group for api/v1/*
func RegisterRouterAPIV1(router *gin.RouterGroup, db *gorm.DB) {
	userAPI := NewUserAPI(db, true)
	router.GET("/users", userAPI.GetAllUser)
	router.GET("/users/:name", userAPI.GetUser)

	prodAPI := NewProductAPI(db, true)
	router.GET("/products", prodAPI.FindAll)
	router.GET("/products/:id", prodAPI.GetProduct)
	router.DELETE("/products/:id", prodAPI.DeleteProduct)
}
