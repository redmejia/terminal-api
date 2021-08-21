package main

import (
	"log"

	"github.com/redmejia/terminal/driver"
)

func main() {
	db, err := driver.Conn()

	if err != nil {
		log.Println("ERROR ", err)
	}
	defer db.Close()

	// testing database connection
	err = db.Ping()
	if err != nil {
		log.Println("ERROR PING ", err)
	}
	log.Println("OK PASS ")
}
