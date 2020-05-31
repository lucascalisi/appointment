package mysql

import (
	"database/sql"
	"fmt"

	rec "github.com/appointment/resources"
)

func (db *DB) GetProfessionalByID(id int64) (rec.Professional, error) {
	result := rec.Professional{}
	err := db.QueryRow("SELECT idUser, dni, name, sex, doctorNumber, birthDay FROM professionals WHERe idUser = ?", id).Scan(&result.ID, &result.DNI, &result.Name, &result.Sex, &result.DoctorNumber, &result.BirthDay)
	if err != nil {
		if err == sql.ErrNoRows {
			return rec.Professional{}, rec.NotFound
		}
		return rec.Professional{}, rec.NewStorageError(fmt.Sprintf("could not get user info  : %v", err))
	}

	return result, nil
}
