package main

import (
	"database/sql"
	"gts-dry/config"
	"gts-dry/router"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	config.Load()

	dbURL := config.GetDBURL()
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := router.SetupRouter(db)

	serverPort := config.GetServerPort()
	app.Listen(":" + serverPort)
}
