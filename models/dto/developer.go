package dto

import "time"

type Developer struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DeveloperDeleted struct {
	Name      string
	DeletedAt time.Time
}
