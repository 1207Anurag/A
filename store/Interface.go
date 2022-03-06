package store

import (
	"Hospital/model"
)

type Patient_er interface {
	Get(id int) (*model.Patient, error)
	GetAll() (model.Patient, error)
	Create(patient model.Patient) error
	Update(id int, Name string, desc string)
}
