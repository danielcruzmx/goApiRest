package database

import "database/sql"

func InitDB() *sql.DB {
	strCnn := "root:admin@tcp(localhost:3306)/northwind"
	dbCnn, err := sql.Open("mysql", strCnn)

	if err != nil {
		panic(err.Error())
	}

	return dbCnn
}
