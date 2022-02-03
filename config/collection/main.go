package collection

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"product-api/app/controllers/product"
)

func MainRouter(db *gorm.DB, main *gin.RouterGroup) {
	ProductCtrl := product.ProductController(db)
	main.GET("/product", ProductCtrl.GetProducts)
	main.POST("/product", ProductCtrl.AddProduct)
}