package usecase

import (
	"github.com/itsapep/golang-with-mongodb/model"
	"github.com/itsapep/golang-with-mongodb/repository"
	"github.com/itsapep/golang-with-mongodb/utils"
)

type FindAllProductUsecase interface {
	FindAllProduct(page int64, limit int64) ([]model.Product, error)
}

type findAllProductUsecase struct {
	repo repository.ProductRepository
}

// Register implements FindAllProductUsecase
func (f *findAllProductUsecase) FindAllProduct(page int64, limit int64) ([]model.Product, error) {
	productsInterface, err := f.repo.Retrieve(page, limit)
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
