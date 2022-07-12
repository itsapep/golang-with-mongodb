package usecase

import (
	"github.com/itsapep/golang-with-mongodb/model"
	"github.com/itsapep/golang-with-mongodb/repository"
)

type UpdateProductUsecase interface {
	UpdateProductById(id string, updatedProduct *model.Product) error
}

type updateProductUsecase struct {
	repo repository.ProductRepository
}

// Register implements UpdateProductUsecase
func (u *updateProductUsecase) UpdateProductById(id string, updatedProduct *model.Product) error {
	err := u.repo.Update(id, updatedProduct)
	if err != nil {
		return err
	}
	return nil
}

func NewUpdateProductUsecase(repo repository.ProductRepository) UpdateProductUsecase {
	return &updateProductUsecase{
		repo: repo,
	}
}
