package routes

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/redmejia/terminal/driver"
	"github.com/redmejia/terminal/handlers"
	"github.com/redmejia/terminal/middleware"
)

func Routes(db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	dbConn := driver.NewDBRepo(db)

	handler := handlers.NewHandler(errLog, infoLog, dbConn)

	middle := middleware.NewMiddleware(infoLog, errLog)

	mux.HandleFunc("/", handler.HandleAuth)

	mux.HandleFunc("/project", handler.HandelProject)
	mux.HandleFunc("/project/like", handler.HandleLike)

	mux.HandleFunc("/project/comment", handler.HandleComment)

	return middle.JsonFormat(middle.Header(middle.RequestLogger(mux)))
}
