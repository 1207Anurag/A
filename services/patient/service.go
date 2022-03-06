package patient

import (
	"Hospital/model"
	"Hospital/store"
	"errors"
)

type Patients struct {
	PatientS store.Patient_er
}

//service layer is communicating with the database layer

func New(c store.Patient_er) Patients {
	return Patients{PatientS: c}
}

func (pat Patients) GetPatientById(id int) (model.Patient, error) {
	if id == 0 {
		return model.Patient{}, errors.New("Invalid Id")
	}
	c, err := pat.PatientS.Get(id)
	if err != nil {
		return model.Patient{}, err
	}
	return pat, nil

}

func (pat Patients) PostPatient(patient model.Patient) (model.Patient, error) {
	c, err := ValidatePatientDetail(patient)
	if err != nil {
		return model.Patient{}, err
	}
	_, err = pat.PatientS.PostPatient(patient)
	if err != nil {
		return model.Patient{}, err
	}
	return c, nil
}

func ValidatePatientDetail(patient models.PatientStructure) (models.PatientStructure, error) {
	if len(patient.Data.Patient.Name) < 1 || patient.Data.Patient.Name == " " {
		return models.PatientStructure{}, nil
	}
	if len(patient.Data.Patient.Phone) < 10 || patient.Data.Patient.Phone == " " {
		return models.PatientStructure{}, nil
	}
	return patient, nil
}
