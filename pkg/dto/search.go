package dto

type ProductCount struct {
	Count					int  				`json:"count"`
	Preferences		Preferences	`json:"preferences"`     // see preference.go
//	Sort          string      `json:"sort"`
}
