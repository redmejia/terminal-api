package handlers

import (
	"log"

	"github.com/redmejia/terminal/driver"
)

// Handler
type Handler struct {
	ErrorLog, InfoLog *log.Logger
	DB                driver.IDatabaseRepo
}

// NewHadler
func NewHandler(errLog, infoLog *log.Logger, DB driver.IDatabaseRepo) *Handler {
	return &Handler{
		ErrorLog: errLog,
		InfoLog:  infoLog,
		DB:       DB,
	}
}
