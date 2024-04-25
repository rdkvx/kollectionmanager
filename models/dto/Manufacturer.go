package dto

import (
	"time"
)

type ManufacturerGet struct {
	Name     string
	Founded  time.Time
}

type ManufacturerPost struct {
	Name     string
	Founded  string
}