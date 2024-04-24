package dto

import (
	"time"
)

type Game struct {
	Name        string
	ConsoleID   uint
	DeveloperID uint
	ReleaseDate time.Time
	BoughtDate  time.Time
}
