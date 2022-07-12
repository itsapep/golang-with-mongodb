package usecase

import (
	"github.com/itsapep/golang-with-mongodb/model"
	"github.com/itsapep/golang-with-mongodb/repository"
)

type ProductRegistrationUsecase interface {
	Register(newProduct *model.Product) error
}

type productRegistrationUsecase struct {
	repo repository.ProductRepository
}

// Register implements ProductRegistrationUsecase
func (p *productRegistrationUsecase) Register(newProduct *model.Product) error {
	return p.repo.Add(newProduct)
}

func NewProductRegistrationUsecase(repo repository.ProductRepository) ProductRegistrationUsecase {
	return &productRegistrationUsecase{
		repo: repo,
	}
}
