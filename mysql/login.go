package mysql

import (
	"database/sql"
	"fmt"

	rec "github.com/appointment/resources"
)

func (db *DB) Login(cred rec.LoginAuth) error {
	var count int
	err := db.QueryRow("SELECT 1 FROM users WHERE email = ? AND password = ? AND status = 'enabled'", cred.User, cred.Password).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return rec.AnauthorizedUser
		}
		return rec.NewStorageError(fmt.Sprintf("could not get user info  : %v", err))
	}

	return nil
}
