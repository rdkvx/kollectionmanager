package models

import (
	"time"

	"gorm.io/gorm"
)

type Console struct {
	gorm.Model
	Name           string `gorm:"not null;unique"`
	ManufacturerID uint   `gorm:"not null;"`
	DtRelease      time.Time
	DtPurchase     time.Time
	Owned          bool   `gorm:"not null;"`
	Deleted        bool   `gorm:"not null;"`
	Games          []Game `gorm:"many2many:game_consoles;"`
}
