package usecase

import (
	"github.com/itsapep/golang-with-mongodb/model"
	"github.com/itsapep/golang-with-mongodb/repository"
	"github.com/itsapep/golang-with-mongodb/utils"
)

type FindAllProductUsecase interface {
	FindAllProduct() ([]model.Product, error)
}

type findAllProductUsecase struct {
	repo repository.ProductRepository
}

// Register implements FindAllProductUsecase
func (f *findAllProductUsecase) FindAllProduct() ([]model.Product, error) {
	productsInterface, err := f.repo.Retrieve()
	if err != nil {
		return nil, err
	}
	products, ok := productsInterface.([]model.Product)
	if ok {
		return products, nil
	}
	return nil, utils.NewCastInterfaceError("product")

}

func NewFindAllProductUsecase(repo repository.ProductRepository) FindAllProductUsecase {
	return &findAllProductUsecase{
		repo: repo,
	}
}
