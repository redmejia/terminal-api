package main

import (
	"log"
	"net/http"

	"github.com/redmejia/terminal/driver"
	"github.com/redmejia/terminal/routes"
)

func main() {
	db, err := driver.Conn()

	if err != nil {
		log.Println("ERROR ", err)
	}
	defer db.Close()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes.Routes(db),
	}

	log.Println("Server is running at :8080")
	log.Fatal(srv.ListenAndServe())

}
