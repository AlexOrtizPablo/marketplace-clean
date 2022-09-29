package in_memory

import "marketplace-clean/domain"

type UserRepositoryInMemory struct {
}

func (r *UserRepositoryInMemory) GetUserByMailAndPassword(mail string, password string) domain.User {
	return domain.User{}
}

func (r *UserRepositoryInMemory) SaveUser(created domain.User) {

}

func (r *UserRepositoryInMemory) GetUserById(id string) *domain.User {
	return &domain.User{}
}

func (r *UserRepositoryInMemory) UpdateUser(user *domain.User) {

}
