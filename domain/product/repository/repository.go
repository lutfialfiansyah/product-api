package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"product-api/domain/product/model"
)

type RepositoryInterface interface {
	GetProducts(sortBy, sorting string) ([]model.Product, error)
	GetCurrencyByID(id int64) (model.Product, error)
	GetByName(name string) (model.Product, error)
	Store(req model.RequestProduct) (model.Product, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) RepositoryInterface {
	return &repository{
		DB: DB,
	}
}

func (r *repository) GetProducts(sortBy, sorting string) ([]model.Product, error) {
	var products []model.Product
	if err := r.DB.Table("products").
		Select("*").
		Order(fmt.Sprintf("%v %v", sortBy, sorting)).
		Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) GetCurrencyByID(id int64) (model.Product, error) {
	var product model.Product
	if err := r.DB.Table("products").
		Select("*").
		Where("id = ?", id).
		Find(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) GetByName(name string) (model.Product, error) {
	var product model.Product
	if err := r.DB.Table("products").
		Select("*").
		Where("name iLIKE ?", `%`+name+`%`).
		Find(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, nil
		}
		return product, err
	}

	return product, nil
}

func (r *repository) Store(req model.RequestProduct) (model.Product, error) {
	var product model.Product
	if err := r.DB.Table("products").
		Create(&req).Error; err != nil {
		return product, err
	}

	// CALLBACK
	product, err := r.GetCurrencyByID(req.ID)
	if err != nil {
		return product, err
	}

	return product, nil
}