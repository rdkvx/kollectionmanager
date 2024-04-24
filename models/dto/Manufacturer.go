package dto

import (
	"time"
)

type Manufacturer struct {
	ID       uint
	Name     string
	Founded  time.Time
	Consoles []Console
}
