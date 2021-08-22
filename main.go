package main

import (
	"log"

	"github.com/redmejia/terminal/driver"
	"github.com/redmejia/terminal/handlers"
)

func main() {
	db, err := driver.Conn()

	if err != nil {
		log.Println("ERROR ", err)
	}
	defer db.Close()

	dbConn := driver.DbRepo{Conn: db}

	var hand handlers.Handler
	hand.DB = &dbConn
	hand.HandleAuth()

}
