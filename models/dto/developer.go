package dto

import "time"

type Developer struct {
	Name string
}

type DeveloperDeleted struct {
	ID        uint
	Name      string
	DeletedAt time.Time
}
