package repositories

type Repo interface {
	FindById(id string) RepositoryResult
	Create(interface{}) RepositoryResult
	FindAll() 			RepositoryResult
	Update(interface{}) RepositoryResult
	Delete(id string) 	RepositoryResult
}
