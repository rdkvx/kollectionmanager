package models

import (
	"time"

	"gorm.io/gorm"
)

type Manufacturer struct {
	gorm.Model
	ID       uint
	Name     string
	Founded  time.Time
	Consoles []Console
}
