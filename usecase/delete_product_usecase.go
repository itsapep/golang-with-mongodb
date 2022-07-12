package usecase

import (
	"github.com/itsapep/golang-with-mongodb/repository"
)

type DeleteProductUsecase interface {
	DeleteProductById(id string) error
}

type deleteProductUsecase struct {
	get  GetProductByIdUsecase
	repo repository.ProductRepository
}

// Register implements DeleteProductUsecase
func (d *deleteProductUsecase) DeleteProductById(id string) error {
	product, err := d.get.GetProductById(id)
	if err != nil {
		return err
	}
	err = d.repo.Delete(&product)
	if err != nil {
		return err
	}
	return nil
}

func NewDeleteProductUsecase(repo repository.ProductRepository) DeleteProductUsecase {
	return &deleteProductUsecase{
		repo: repo,
	}
}
