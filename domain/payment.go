package domain

import (
	"errors"
	"time"
)

type Payment struct {
	Id        string
	Payer     User
	Amount    int64
	Status    string
	CreatedAt time.Time
}

func CreatePayment(payment Payment) Payment {
	payment.Status = "pending"
	payment.CreatedAt = time.Time{}
	return payment
}

func (p *Payment) IsValid() bool {
	if p.Id == "" {
		return false
	}
	if !p.Payer.IsValid() {
		return false
	}
	if p.Amount <= 0 {
		return false
	}
	if p.Status == "" {
		return false
	}

	return true
}

func (p *Payment) Approve() error {
	if p.Status == "approved" || p.Status == "canceled" {
		return errors.New("cannot approve payment")
	}

	return nil
}

func (p *Payment) Cancel() error {
	if p.Status == "canceled" || p.Status == "approved" {
		return errors.New("cannot approve payment")
	}

	return nil
}
