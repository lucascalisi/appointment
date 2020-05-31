package mysql

import (
	"database/sql"
	"fmt"

	rec "github.com/appointment/resources"
)

func (db *DB) GetPatientByID(id int64) (rec.Patient, error) {
	result := rec.Patient{}
	err := db.QueryRow("SELECT idUser, dni, name, sex, birthDay FROM patients WHERe idUser = ?", id).Scan(&result.ID, &result.DNI, &result.Name, &result.Sex, &result.BirthDay)
	if err != nil {
		if err == sql.ErrNoRows {
			return rec.Patient{}, rec.NotFound
		}
		return rec.Patient{}, rec.NewStorageError(fmt.Sprintf("could not get user info  : %v", err))
	}

	return result, nil
}
