package routes

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/redmejia/terminal/driver"
	"github.com/redmejia/terminal/handlers"
)

func Routes(db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	dbConn := driver.NewDBRepo(db)

	var handler = handlers.Handler{
		DB:         dbConn,
		SuccessLog: log.New(os.Stdout, "SUCCESS\t", log.Ldate|log.Ltime),
		InfoLog:    log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLog:   log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	mux.HandleFunc("/", handler.HandleAuth)

	return mux
}
