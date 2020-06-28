package mysql

import (
	"fmt"

	rec "github.com/appointment/resources"
)

func (db *DB) GetQueue() (rec.Queue, error) {
	queue := rec.Queue{}
	rows, err := db.Query("SELECT id, idPatient, idSpecialityDetail, status FROM queue ORDER BY id DESC")
	if err != nil {
		return rec.Queue{}, rec.NewStorageError(fmt.Sprintf("could not get queue: %v", err))
	}

	defer rows.Close()

	for rows.Next() {
		i := rec.QueueItem{}
		err := rows.Scan(&i.ID, &i.IDPatient, &i.IDSpecialty, &i.Status)
		if err != nil {
			return rec.Queue{}, rec.StorageError{
				Description: fmt.Sprintf("could not scan queue: %v", err),
			}
		}
		queue.Items = append(queue.Items, i)
	}

	return queue, nil
}

func (db *DB) AddItem(item rec.QueueItem) error {
	_, err := db.Exec(`INSERT INTO queue (idPatient, idSpecialityDetail, status) VALUES(?, ?, ?)`, item.IDPatient, item.IDSpecialty, "waiting")
	if err != nil {
		return rec.NewStorageError(fmt.Sprintf("could not add patient to queue: %v", err))
	}

	return nil
}
