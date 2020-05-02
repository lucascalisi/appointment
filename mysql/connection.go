package mysql

import (
	"database/sql"
	"fmt"

	app_models "github.com/appointment/app_models"
	"github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func NewDB(dbUser, dbPassword, dbHost string, dbPort int, dbName string, debug bool) (*DB, error) {
	dbConStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dbConStr)
	if err != nil {
		return nil, app_models.NewStorageError(fmt.Sprintf("could not connect to the database: %v", err))
	}

	err = db.Ping()
	if err != nil {
		return nil, app_models.NewStorageError(fmt.Sprintf("could not establish connection with the database: %v", err))
	}

	mysqlDB := DB{db}
	return &mysqlDB, nil
}

func SetConnectionParameters(db *DB) {
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(10)
}

// Associates an existing vulnerability to a report. Assumes report and vulnerability exist.
// Returns report.StorageError in case of database error.
func (db *DB) AssociateVulnerability(repID, vulID, hostID int64) error {
	_, err := db.Exec("INSERT INTO reports_vulnerabilities_hosts (report_id, vulnerability_id, host_id) VALUES (?, ?, ?)", repID, vulID, hostID)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number != 1062 { // Check that the error is not a duplicated FK error
				return app_models.NewStorageError(fmt.Sprintf("could not add vulneerability to report: %v", err))
			}
		}
	}

	return nil
}
