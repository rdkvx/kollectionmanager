package models

import (
	"time"
)

type Game struct {
	Id            string    `json:"id", omitempty`
	Name          string    `json:"name", omitempty`
	IdConsole     string    `json:"id_console", omitempty`
	IdDeveloper   string    `json:"id_developer", omitempty`
	DtReleaseDate time.Time `json:"dt_release_date", omitempty`
	DtBought      time.Time `json:"dt_buy", omitempty`
}
