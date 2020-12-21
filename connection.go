package main

import(
	"database/sql"
	_ "github.com/lib/pq"

)
//getConnection obtener conexión BD

func getConnection() *sql.DB {
	dsn := "postgres://golag:123@127.0.0.1:5432/gocrud?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err !={
		log
	}
	
}