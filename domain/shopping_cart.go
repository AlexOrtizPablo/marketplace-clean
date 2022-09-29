package domain

import (
	"time"
)

type ShoppingCart struct {
	Id         string
	Buyer      User
	Products   map[Product]int
	TotalPrice int64
	UpdatedAt  time.Time
}
