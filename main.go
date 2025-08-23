package main

import (
	"database/sql"
	"tugas-13/controllers"
	"tugas-13/routers"

	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "sadam"
// 	dbname   = "bioskop"
// )

var db *sql.DB
var err error

func main() {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	os.Getenv("PGHOST"),
	// 	5432,
	// 	os.Getenv("PGUSER"),
	// 	os.Getenv("PGPASSWORD"),
	// 	os.Getenv("PGDATABASE"),
	// )
	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// db, err = sql.Open("postgres", psqlInfo)

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	controllers.DB = db

	fmt.Println("Successfully Connected do database")

	routers.StartServer().Run(os.Getenv("PORT"))

}
