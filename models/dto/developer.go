package dto

import "gorm.io/gorm"

type Developer struct {
	gorm.Model
	ID   uint 
	Name string 
	Games []Game 
}
