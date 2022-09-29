package domain

import (
	"errors"
	"time"
)

type User struct {
	Id        string
	Name      string
	LastName  string
	Password  string
	Mail      string
	Address   Address
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateUser(user User) (User, error) {
	if !user.IsValid() {
		return user, errors.New("user not valid")
	}
	user.CreatedAt = time.Time{}
	return user, nil
}

func UpdateUser(user User) (User, error) {
	if !user.IsValid() {
		return user, errors.New("user not valid")
	}
	user.UpdatedAt = time.Time{}
	return user, nil
}

func (u *User) IsValid() bool {
	if u.Id == "" {
		return false
	}
	if u.Name == "" {
		return false
	}
	if u.LastName == "" {
		return false
	}
	if u.Password == "" {
		return false
	}
	if u.Mail == "" {
		return false
	}
	if !u.Address.IsValid() {
		return false
	}

	return true
}
