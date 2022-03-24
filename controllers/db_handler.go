package controllers

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_latihan_pbp?charset=utf8mb4&parseTime=True&loc=Local")
	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_latihan_pbp")

	if err != nil {
		log.Fatal(err)
	}
	return db
}
