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

	dbConn := driver.NewDBRepo(db)

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	var handler = handlers.Handler{
		DB:       dbConn,
		InfoLog:  infoLog,
		ErrorLog: errLog,
	}

	var middle = middleware.Middleware{
		InfoLog:  infoLog,
		ErrorLog: errLog,
	}

	mux.HandleFunc("/", handler.HandleAuth)

	mux.HandleFunc("/project", handler.HandelProject)
	mux.HandleFunc("/project/like", handler.HandleLike)

	mux.HandleFunc("/project/comment", handler.HandleComment)

	return middle.Loggers(mux)
}
