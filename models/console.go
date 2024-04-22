package models

import "time"

type Console struct {
	Id             string    `json:"id", omitempty"`
	Name           string    `json:"name", omitempty"`
	IdManufacturer  string    `json:"id_manufacturer", omitempty`
	ReleaseDate  time.Time `json:"dt_release_date", omitempty`
	PurchaseDate time.Time `json:"dt_purchase_date", omitempty`
	Owned          bool      `json:"owned", omitempty`
}
