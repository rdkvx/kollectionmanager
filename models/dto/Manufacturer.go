package dto

import (
	"time"
)

type Manufacturer struct {
	Name     string
	Founded  time.Time
	Consoles []Console
}
