package domain

import (
	"time"
)

type Question struct {
	Id           string
	Product      Product
	UserWhoAsked User
	Description  string
	Answer       string
	CreatedAt    time.Time
	AnsweredAt   time.Time
}
