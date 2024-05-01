package dto

import (
	"time"
)

type GameGet struct {
	ID          uint
	Name        string
	ConsoleID   uint
	DeveloperID uint
	ReleaseDate time.Time
	BoughtDate  time.Time
	Owned       bool
}

type GamePost struct {
	Name        string
	ConsoleID   uint
	DeveloperID uint
	ReleaseDate string
	BoughtDate  string
	Owned       bool
}
