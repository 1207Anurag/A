package patient

import (
	"Hospital/model"
	"database/sql"
	"fmt"
)

type PatientStorer struct {
	db *sql.DB
}

func New(db *sql.DB) PatientStorer {
	return PatientStorer{db: db}
}

func (a PatientStorer) Get(id int) (*model.Patient, error) {
	var (
		rows *sql.Rows
		err  error
	)

	rows, err = a.db.Query("SELECT * FROM patient WHERE id=?", id)

	defer rows.Close()
	var k model.Patient
	//var pat []model.Patient
	for rows.Next() {
		err = rows.Scan(&k.Id, &k.Name, &k.Phone, &k.Discharge,
			&k.BloodGroup, &k.Description)
		if err != nil {
			return nil, err
		}
	}
	return &k, nil
}
func (a PatientStorer) GetAll() (model.Patient, error) {
	var (
		rows *sql.Rows
		err  error
	)
	rows, err = a.db.Query("SELECT * FROM patient")
	defer rows.Close()
	var k model.Patient
	err = rows.Scan(&k.Id, &k.Name, &k.Phone, &k.Discharge,
		&k.BloodGroup, &k.Description)
	if err != nil {
		return model.Patient{}, nil
	}
	return k, nil
}
func (a PatientStorer) Create(p model.Patient) error {
	query := "INSERT INTO patient (Id,Name,Phone,Discharge,BloodGroup,Description) values(?,?,?,?,?,?)"
	res, err := a.db.Exec(query, p.Id, p.Name, p.Discharge, p.BloodGroup, p.Description)
	if err != nil {
		return err
	}
	lastId, er := res.LastInsertId()
	if er != nil {
		return er
	}
	fmt.Printf("Last inseted id %d\n", lastId)
	return err
}

func (a PatientStorer) Update(id int, Name string, desc string) error {
	_, err := a.db.Exec("UPDATE patient SET name=? desc=? WHERE id=?", id, Name, desc)
	if err != nil {
		return err
	}
	return nil
}
