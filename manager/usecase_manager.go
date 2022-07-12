package manager

import "github.com/itsapep/golang-with-mongodb/usecase"

type UsecaseManager interface {
	ProductRegistrationUsecase() usecase.ProductRegistrationUsecase
	FindAllProductUsecase() usecase.FindAllProductUsecase
	UpdateProductUsecase() usecase.UpdateProductUsecase
	DeleteProductUsecase() usecase.DeleteProductUsecase
	GetProductByIdUsecase() usecase.GetProductByIdUsecase
	GetProductByCategoryUsecase() usecase.GetProductByCategoryUsecase
}

type usecaseManager struct {
	repoManager RepositoryManager
}

// DeleteProductUsecase implements UsecaseManager
func (u *usecaseManager) DeleteProductUsecase() usecase.DeleteProductUsecase {
	return usecase.NewDeleteProductUsecase(u.repoManager.ProductRepo())
}

// FindAllProductUsecase implements UsecaseManager
func (u *usecaseManager) FindAllProductUsecase() usecase.FindAllProductUsecase {
	return usecase.NewFindAllProductUsecase(u.repoManager.ProductRepo())
}

// GetProductByCategoryUsecase implements UsecaseManager
func (u *usecaseManager) GetProductByCategoryUsecase() usecase.GetProductByCategoryUsecase {
	return usecase.NewGetProductByCategoryUsecase(u.repoManager.ProductRepo())
}

// GetProductByIdUsecase implements UsecaseManager
func (u *usecaseManager) GetProductByIdUsecase() usecase.GetProductByIdUsecase {
	return usecase.NewGetProductByIdUsecase(u.repoManager.ProductRepo())
}

// UpdateProductUsecase implements UsecaseManager
func (u *usecaseManager) UpdateProductUsecase() usecase.UpdateProductUsecase {
	return usecase.NewUpdateProductUsecase(u.repoManager.ProductRepo())
}

// ProductRegistrationUsecase implements UsecaseManager
func (u *usecaseManager) ProductRegistrationUsecase() usecase.ProductRegistrationUsecase {
	return usecase.NewProductRegistrationUsecase(u.repoManager.ProductRepo())
}

func NewUsecaseManager(repoManager RepositoryManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
