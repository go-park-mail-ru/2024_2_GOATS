package postgres

import (
	"fmt"
)

type connectData struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
}

func getConnectData() *connectData {
	connData := &connectData{
		//host:     os.Getenv("DB_HOST"),
		//port:     os.Getenv("DB_PORT"),
		//user:     os.Getenv("DB_USER"),
		//password: os.Getenv("DB_PASS"),
		//dbName:   os.Getenv("DB_NAME"),

		//DB_PASS = 1111
		//DB_USER = postgres
		//DB_NAME = Netflix
		//DB_HOST = localhost
		//DB_PORT = 5432

		host:     "localhost",
		port:     "5432",
		user:     "postgres",
		password: "1111",
		dbName:   "Netflix",
	}
	fmt.Println("host = ", connData.host)
	return connData
}
