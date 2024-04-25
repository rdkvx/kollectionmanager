package models

import "gorm.io/gorm"

type Developer struct {
	gorm.Model
	Name string  `gorm:"not null;unique"`
	Deleted bool
	Games []Game 
}
