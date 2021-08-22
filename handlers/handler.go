package handlers

import (
	"log"

	"github.com/redmejia/terminal/driver"
)

type Handler struct {
	ErrorLog, InfoLog, SuccessLog *log.Logger
	DB                            driver.IDatabaseRepo
}
