package product

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"product-api/domain/product"
	"product-api/domain/product/model"

	"product-api/domain/product/repository"
	"product-api/lib/response"
)

type Controller struct {
	ProductService product.ServiceInterface
}

func ProductController(db *gorm.DB) *Controller {
	return &Controller{
		ProductService: product.NewService(repository.NewRepository(db)),
	}
}

func (av *Controller) GetProducts(context *gin.Context) {
	sorting := context.DefaultQuery("sort", "ASC")
	sortBy := context.DefaultQuery("sort_by", "name")
	resBody, errStatus, err := av.ProductService.GetProducts(sortBy, sorting)
	if err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}
	response.Json(context, http.StatusOK, resBody)
}

func (av *Controller) AddProduct(ctx *gin.Context) {
	var req model.RequestProduct
	if err := ctx.Bind(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// REQUEST VALIDATION
	valid, err := govalidator.ValidateStruct(req)
	if err != nil {
		println("error: " + err.Error())
		response.Error(ctx, http.StatusBadRequest, "Invalid Input : "+err.Error())
		return
	}
	if !valid {
		response.Error(ctx, http.StatusBadRequest, "Invalid Input : "+err.Error())
		return
	}

	data, errStatus, err := av.ProductService.AddProduct(req)
	if err != nil {
		response.Error(ctx, errStatus, err.Error())
		return
	}
	response.Json(ctx, http.StatusOK, data)
}
