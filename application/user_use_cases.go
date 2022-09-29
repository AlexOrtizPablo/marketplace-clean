package application

import (
	"errors"
	"marketplace-clean/application/interfaces"
	"marketplace-clean/domain"
)

type UserUseCases struct {
	UserRepository interfaces.UserRepository
}

func (u *UserUseCases) RegisterNew(user domain.User) error {
	userCreated, err := domain.CreateUser(user)
	if err != nil {
		return errors.New("error creating user")
	}

	u.UserRepository.SaveUser(userCreated)

	return nil
}

func (u *UserUseCases) UpdateUserAddress(userId string, address domain.Address) error {
	user := u.UserRepository.GetUserById(userId)

	user.Address = address
	if !user.IsValid() {
		return errors.New("user not valid")
	}

	u.UserRepository.UpdateUser(user)

	return nil
}

func (u *UserUseCases) Login(mail string, password string) (domain.User, error) {
	user := u.UserRepository.GetUserByMailAndPassword(mail, password)

	if !user.IsValid() {
		return domain.User{}, errors.New("incorrect mail or password")
	}

	return user, nil
}

func (u *UserUseCases) PublishNewProduct(user domain.User, product domain.Product) error {
	return nil
}

func (u *UserUseCases) PublishNewQuestion(user domain.User, product domain.Product, question domain.Question) error {
	return nil
}

func (u *UserUseCases) PublishNewRating(user domain.User, product domain.Product, rating domain.Rating) error {
	return nil
}

func (u *UserUseCases) CreateShoppingCart(user domain.User, cart domain.ShoppingCart) error {
	return nil
}

func (u *UserUseCases) AddProductToUserShoppingCart(user domain.User, cart domain.ShoppingCart) error {
	return nil
}

func (u *UserUseCases) BuyCart(user domain.User) error {
	return nil
}

func (u *UserUseCases) ReceiveShipping(user domain.User) error {
	return nil
}
