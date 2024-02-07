package store

import (
	"database/sql"
	"fmt"
)

func GetDBInstance() (*sql.DB, error) {
	db, errr := ConnectDB()
	if errr != nil {
		fmt.Println("Error ", errr)
		return db, errr
	}
	return db, nil
}
func ConnectDB() (*sql.DB, error) {
	connStr := "host=localhost port=8080 user=postgres password=Dhruv12345 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
