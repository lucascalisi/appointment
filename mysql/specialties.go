package mysql

import (
	"fmt"

	rec "github.com/appointment/resources"
)

func (db *DB) Search(params rec.SpecialtySearcher) ([]rec.Specialty, error) {
	params.Category = "%" + params.Category + "%"
	params.SubCategory = "%" + params.SubCategory + "%"

	query := `SELECT s.id, s.name, sd.id, sd.name
	FROM specialties s
	INNER JOIN specialityDetails sd
		ON s.id = sd.idSpeciality
	WHERE s.name LIKE ?
	AND sd.name LIKE ?
	ORDER BY s.id`

	rows, err := db.Query(query, params.Category, params.SubCategory)
	if err != nil {
		return nil, rec.NewStorageError(fmt.Sprintf("could not search specialties info  : %v", err))
	}

	defer rows.Close()

	specialties := []rec.Specialty{}
	specialtiDetails := []rec.SpecialtyDetails{}
	sNext := rec.Specialty{}
	s := rec.Specialty{}
	for rows.Next() {
		var idSubCategory *int64
		var subCategory *string

		err := rows.Scan(&sNext.ID, &sNext.Category, &idSubCategory, &subCategory)
		if err != nil {
			return nil, rec.StorageError{
				Description: fmt.Sprintf("could not search specialties: %v", err),
			}
		}

		if s.ID != sNext.ID {
			s.SubCategories = specialtiDetails
			specialties = append(specialties, s)
			specialtiDetails = nil
		}

		if idSubCategory != nil {
			specialtiDetails = append(specialtiDetails, rec.SpecialtyDetails{
				ID:          *idSubCategory,
				SubCategory: *subCategory,
			})
		}
		s = sNext
	}

	s.SubCategories = specialtiDetails
	specialties = append(specialties, s)

	return specialties, nil
}

func (db *DB) GetSpecialtyByDetail(id int64) (rec.Specialty, error) {
	query := `SELECT s.id, s.name, sd.id, sd.name 
	FROM specialityDetails sd
	INNER JOIN specialties s
		ON s.id = sd.idSpeciality
	WHERE sd.id = ?`

	var idSubCategory *int64
	var subCategory *string
	specialty := rec.Specialty{}

	err := db.QueryRow(query, id).Scan(&specialty.ID, &specialty.Category, &idSubCategory, &subCategory)
	if err != nil {
		return rec.Specialty{}, rec.NewStorageError(fmt.Sprintf("could not get specialties by details: %v", err))
	}

	specialtyDetail := rec.SpecialtyDetails{}
	if idSubCategory != nil {
		specialtyDetail.ID = *idSubCategory

	}
	if subCategory != nil {
		specialtyDetail.SubCategory = *subCategory

	}

	specialty.SubCategories = append(specialty.SubCategories, specialtyDetail)

	return specialty, nil
}
