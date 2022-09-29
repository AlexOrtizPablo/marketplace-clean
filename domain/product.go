package domain

import (
	"time"
)

type Product struct {
	Id          string
	Name        string
	Price       string
	Description string
	Seller      User
	Stock       int
	CreatedAt   time.Time
}
