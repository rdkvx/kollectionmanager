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
	Deleted bool
	Consoles []Console
}
