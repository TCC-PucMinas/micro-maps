package models

type Map struct {
	Street   string `json:"street"`
	District string `json:"district"`
	City     string `json:"city"`
	Country  string `json:"country"`
	State    string `json:"state"`
	Number   string `json:"number"`
	ZipCode  string `json:"zipCode"`
	Lat      string `json:"lat"`
	Lng      string `json:"lng"`
}
