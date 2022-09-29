package interfaces

import (
	"marketplace-clean/domain"
)

type UserRepository interface {
	GetUserByMailAndPassword(mail string, password string) domain.User
	SaveUser(created domain.User)
	GetUserById(id string) *domain.User
	UpdateUser(user *domain.User)
}
