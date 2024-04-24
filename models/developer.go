package models

import "gorm.io/gorm"

type Developer struct {
	gorm.Model
	ID   uint 
	Name string 
	Deleted bool
	Games []Game 
}
