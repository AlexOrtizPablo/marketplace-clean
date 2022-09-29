package domain

import (
	"time"
)

type Purchase struct {
	Id        string
	Buyer     User
	Product   Product
	CreatedAt time.Time
}
