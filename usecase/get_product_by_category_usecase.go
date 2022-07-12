package usecase

import (
	"github.com/itsapep/golang-with-mongodb/model"
	"github.com/itsapep/golang-with-mongodb/repository"
	"github.com/itsapep/golang-with-mongodb/utils"
)

type GetProductByCategoryUsecase interface {
	GetProductByCategory(category string) ([]model.Product, error)
}

type getProductByCategoryUsecase struct {
	repo repository.ProductRepository
}

// Register implements GetProductByCategoryUsecase
func (g *getProductByCategoryUsecase) GetProductByCategory(category string) ([]model.Product, error) {
	productInterface, err := g.repo.Get(category)
	if err != nil {
		return nil, err
	}
	product, ok := productInterface.([]model.Product)
	if ok {
		return product, nil
	}
	return nil, utils.NewCastInterfaceError("product")
}

func NewGetProductByCategoryUsecase(repo repository.ProductRepository) GetProductByCategoryUsecase {
	return &getProductByCategoryUsecase{
		repo: repo,
	}
}
