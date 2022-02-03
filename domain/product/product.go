package product

import (
	"errors"
	"net/http"
	"product-api/domain/product/model"
	"product-api/domain/product/repository"
	"product-api/lib/constant"
)

type ServiceInterface interface {
	GetProducts(sortBy, sorting string) (*[]model.Product, int, error)
	AddProduct(product model.RequestProduct) (model.Product, int, error)
}

type service struct {
	Repository repository.RepositoryInterface
}

func NewService(repository repository.RepositoryInterface) ServiceInterface {
	return &service{
		Repository: repository,
	}
}

func (s *service) GetProducts(sortBy, sorting string) (*[]model.Product, int, error) {
	if sortBy == "" {
		sortBy = "name"
	}
	if sorting == "" {
		sorting = "ASC"
	}
	products, err := s.Repository.GetProducts(sortBy, sorting)
	if err != nil {
		return &products, http.StatusInternalServerError, err
	}

	return &products, http.StatusOK, nil
}

func (s *service) AddProduct(req model.RequestProduct) (model.Product, int, error) {
	var store model.Product
	getByName, err := s.Repository.GetByName(req.Name)
	if err != nil {
		return store, http.StatusInternalServerError, err
	}
	if getByName.ID != 0 {
		return store, http.StatusBadRequest, errors.New(constant.AlreadyExists)
	}

	data, err := s.Repository.Store(req)
	if err != nil {
		return store, http.StatusInternalServerError, err
	}

	return data, http.StatusOK, nil
}
