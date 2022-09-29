package domain

import (
	"errors"
	"time"
)

type Shipping struct {
	Id        string
	User      User
	Status    string
	CreatedAt time.Time
}

func CreateShipping(shipping Shipping) Shipping {
	shipping.CreatedAt = time.Time{}
	return shipping
}

func (s *Shipping) IsValid() bool {
	if s.Id == "" {
		return false
	}
	if s.Status == "" {
		return false
	}
	if s.Id == "" {
		return false
	}
	if !s.User.IsValid() {
		return false
	}

	return false
}

func (s *Shipping) ChangeToInProgress() error {
	if s.Status == "in_progress" || s.Status == "delivered" {
		return errors.New("cannot change shipping status to in progress")
	}

	s.Status = "in_progress"
	return nil
}

func (s *Shipping) ChangeToInDelivered() error {
	if s.Status == "pending" || s.Status == "delivered" {
		return errors.New("cannot change shipping status to in progress")
	}

	s.Status = "delivered"
	return nil
}
