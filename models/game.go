package models

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Name        string `gorm:"not null;unique"`
	ConsoleID   uint   `gorm:"not null"`
	DeveloperID uint   `gorm:"not null"`
	ReleaseDate time.Time
	BoughtDate  time.Time
	Deleted     bool `gorm:"not null;"`
	Owned       bool 
}
