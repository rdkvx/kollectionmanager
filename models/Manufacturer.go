package models

import (
    "time"
)

type Manufacturer struct {
	id string `json:"id", omitempty`
	name string `json:"name", omitempty`
	founded time.time `json:"founded", omitempty`
}
