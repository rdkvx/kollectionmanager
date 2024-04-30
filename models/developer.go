package models

import "gorm.io/gorm"

type Developer struct {
	gorm.Model
	Name string  `gorm:"not null;unique"`
	Deleted bool `gorm:"not null;"`
	Games []Game 
}
