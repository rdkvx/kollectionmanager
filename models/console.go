package models

import (
	"time"

	"gorm.io/gorm"
)

type Console struct {
	gorm.Model
	Name           string `gorm:"not null;unique"`
	ManufacturerID uint `gorm:"not null;"`
	ReleaseDate    time.Time
	PurchaseDate   time.Time
	Owned          bool
	Deleted        bool
	Games          []Game `gorm:"many2many:game_consoles;"`
}
