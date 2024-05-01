package dto

import (
	"time"
)

type ConsoleGet struct {
	ID             uint
	Name           string
	ManufacturerID uint
	DtRelease      time.Time
	DtPurchase     time.Time
	Owned          bool
	Games          []GameGet
}

type ConsolePost struct {
	Name           string
	ManufacturerID uint
	DtRelease      string
	DtPurchase     string
	Owned          bool
}
