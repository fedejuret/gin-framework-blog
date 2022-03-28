package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func GetConnection() *sql.DB {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	username := viper.GetString("DB_USERNAME")
	password := viper.GetString("DB_PASSWORD")
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	database := viper.GetString("DB_DATABASE")

	connectionString := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?parseTime=true"

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err.Error())
	}

	return db
}
