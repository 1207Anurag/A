package vehicle

import (
	"VMS/model"
	"VMS/store"
	"database/sql"
	//"errors"
	"fmt"

	"time"
)

type Dbstore struct {
	db *sql.DB
}

func New(db *sql.DB) store.Store {
	return &Dbstore{db: db}
}

func (d *Dbstore) Insert(v model.Vehicles) (int, error) {
	q := "INSERT INTO vehicles VALUES(?,?,?,?,?,?,?,?,?)"
	res, e := d.db.Exec(q, v.Id, v.Model, v.Color, v.NumPlate, time.Now(), time.Now(), time.Now(), v.Name, v.Launched)
	if e != nil {
		return -1, e
	}
	lastId, er := res.LastInsertId()
	if er != nil {
		return -1, e
	}
	fmt.Printf("Last inseted id %d\n", lastId)
	return int(lastId), nil

}

func (d *Dbstore) GetAll(id int) ([]model.Vehicles, error) {
	var v []model.Vehicles
	results, err := d.db.Query("SELECT * FROM vehicles")
	if err != nil {
		return []model.Vehicles{}, err
	}
	defer results.Close()
	for results.Next() {
		var vh model.Vehicles
		err = results.Scan(&vh.Id, &vh.Model, &vh.Color, &vh.NumPlate, &vh.CreatedAt, &vh.UpdatedAt, &vh.DeletedAt, &vh.Name, &vh.Launched)
		if err != nil {
			return []model.Vehicles{}, err
		}
		v = append(v, vh)
	}
	return v, nil
}
func (d *Dbstore) GetById(Id int) (model.Vehicles, error) {
	var vh model.Vehicles
	row := d.db.QueryRow("SELECT * FROM vehicles WHERE id=?", Id)

	err := row.Scan(&vh.Id, &vh.Model, &vh.Color, &vh.NumPlate, &vh.CreatedAt, &vh.UpdatedAt, &vh.DeletedAt, &vh.Name, &vh.Launched)
	if err != nil {

		return vh, err

	}
	return vh, nil
}

func (d *Dbstore) Update(id int, name string) (err error) {
	_, e := d.db.Exec("UPDATE vehicles SET name ? WHERE id=?", id, name)
	if e != nil {

		return e
	}
	return nil
}
func (d *Dbstore) Remove(Id int) error {

	_, e := d.db.Exec("DELETE FROM vehicles  id=?", Id)
	if e != nil {

		return e
	}
	return nil
}
