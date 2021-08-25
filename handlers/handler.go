package handlers

import (
	"log"

	"github.com/redmejia/terminal/driver"
)

type Handler struct {
	ErrorLog, InfoLog *log.Logger
	DB                driver.IDatabaseRepo
}
