package mysql

import (
	"fmt"

	rec "github.com/appointment/resources"
)

func (db *DB) Login(cred rec.LoginAuth) (rec.User, error) {
	query := `SELECT u.id, u.email, u.status, r.id, r.name, r.description
	FROM users u
	INNER JOIN users_roles ur ON u.id = ur.idUser 
	INNER JOIN roles r ON r.id = ur.idRole
	WHERE email = ? AND password = ? AND status = 'enabled'`
	rows, err := db.Query(query, cred.User, cred.Password)
	if err != nil {
		return rec.User{}, rec.NewStorageError(fmt.Sprintf("could not get user info  : %v", err))
	}

	defer rows.Close()

	user := rec.User{}
	roles := []rec.Role{}
	count := 0
	for rows.Next() {
		var roleDescription *string
		count += 1
		r := rec.Role{}
		err := rows.Scan(&user.ID, &user.Email, &user.Status, &r.ID, &r.Name, &roleDescription)
		if err != nil {
			fmt.Println(err)
			return rec.User{}, rec.StorageError{
				Description: fmt.Sprintf("could not scan user: %v", err),
			}
		}

		if roleDescription != nil {
			r.Description = *roleDescription
		}

		roles = append(roles, r)
	}

	if count == 0 {
		return rec.User{}, rec.UnauthorizedUser
	}
	user.Roles = roles

	return user, nil
}
