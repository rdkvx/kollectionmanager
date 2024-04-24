package models

import (
	"time"

	"gorm.io/gorm"
)

type Console struct {
	gorm.Model
	ID             uint     
	Name           string   
	ManufacturerID uint     
	ReleaseDate    time.Time
	PurchaseDate   time.Time
	Owned          bool     
	Games          []Game `gorm:"many2many:game_consoles;"`
}
