package dto

import "time"

type Developer struct {
	Name      string
}

type DeveloperDeleted struct {
	Name      string
	DeletedAt time.Time
}
