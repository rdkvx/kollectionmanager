package dto

import (
	"time"
)

type Game struct {
	ID          uint
	Name        string
	ConsoleID   uint
	DeveloperID uint
	ReleaseDate time.Time
	BoughtDate  time.Time
}
