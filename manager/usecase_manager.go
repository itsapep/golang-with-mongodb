package manager

import "github.com/itsapep/golang-with-mongodb/usecase"

type UsecaseManager interface {
	ProductRegistrationUsecase() usecase.ProductRegistrationUsecase
}

type usecaseManager struct {
	repoManager RepositoryManager
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
