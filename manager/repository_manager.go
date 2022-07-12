package manager

import "github.com/itsapep/golang-with-mongodb/repository"

type RepositoryManager interface {
	ProductRepo() repository.ProductRepository
}

type repositoryManager struct {
	infraManager InfraManager
}

// ProductRepo implements RepositoryManager
func (r repositoryManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(r.infraManager.DBConn())
}

func NewRepositoryManager(infraManager InfraManager) RepositoryManager {
	return repositoryManager{
		infraManager: infraManager,
	}
}
