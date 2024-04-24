package dto

import (
	"time"
)

type Console struct {
	Name           string
	ManufacturerID uint
	ReleaseDate    time.Time
	PurchaseDate   time.Time
	Owned          bool
	Games          []Game
}
