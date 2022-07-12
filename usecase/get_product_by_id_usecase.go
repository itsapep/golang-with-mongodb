package usecase

import (
	"github.com/itsapep/golang-with-mongodb/model"
	"github.com/itsapep/golang-with-mongodb/repository"
	"github.com/itsapep/golang-with-mongodb/utils"
)

type GetProductByIdUsecase interface {
	GetProductById(id string) (model.Product, error)
}

type getProductByIdUsecase struct {
	repo repository.ProductRepository
}

// Register implements GetProductByIdUsecase
func (g *getProductByIdUsecase) GetProductById(id string) (model.Product, error) {
	productInterface, err := g.repo.Get(id)
	if err != nil {
		return model.Product{}, err
	}
	product, ok := productInterface.(model.Product)
	if ok {
		return product, nil
	}
	return model.Product{}, utils.NewCastInterfaceError("product")
}

func NewGetProductByIdUsecase(repo repository.ProductRepository) GetProductByIdUsecase {
	return &getProductByIdUsecase{
		repo: repo,
	}
}
