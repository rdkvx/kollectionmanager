package models

import (
	"time"

	"gorm.io/gorm"
)

type Manufacturer struct {
	gorm.Model
	Name     string `gorm:"not null;unique"`
	Founded  time.Time
	Deleted bool
	Consoles []Console
}
