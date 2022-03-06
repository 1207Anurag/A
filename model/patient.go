package model

type Patient struct {
	Id          int    `json:"id"`
	Name        string `json:"Name"`
	Phone       string `json:"Phone"`
	Discharge   bool   `json:"Discharge"`
	BloodGroup  string `json:"BloodGroup"`
	Description string `json:"Description"`
}
