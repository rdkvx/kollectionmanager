package models

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	ID          uint    
	Name        string    
	ConsoleID   uint    
	DeveloperID uint    
	ReleaseDate time.Time 
	BoughtDate  time.Time 
	Deleted bool
}
