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

	var terminal handlers.Handler
	terminal.DB = dbConn
	terminal.SuccessLog = log.New(os.Stdout, "Success\t", log.Ldate|log.Ltime)

	mux.HandleFunc("/", terminal.HandleAuth)

	return mux
}
