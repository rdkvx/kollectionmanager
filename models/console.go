package models

import "time"

type Console struct {
	Id             string    `json:"id", omitempty"`
	Name           string    `json:"name", omitempty"`
	IdManufacturer  string    `json:"id_manufacturer", omitempty`
	DtReleaseDate  time.Time `json:"dt_release_date", omitempty`
	DtPurchaseDate time.Time `json:"dt_purchase_date", omitempty`
	Owned          bool      `json:"owned", omitempty`
}
