package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:123456789@/golang_gin_posts?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	return db
}
