package domain

import (
	"time"
)

type Rating struct {
	Id           string
	ProductRated Product
	Stars        uint8
	Title        string
	Comment      string
	CreatedAt    time.Time
}
