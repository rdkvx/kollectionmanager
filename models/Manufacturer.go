package models

import (
	"time"
)

type Manufacturer struct {
	Id      string    `json:"id", omitempty`
	Name    string    `json:"name", omitempty`
	Founded time.Time `json:"founded", omitempty`
}
