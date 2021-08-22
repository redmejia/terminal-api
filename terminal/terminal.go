package terminal

import (
	"log"

	"github.com/redmejia/terminal/driver"
)

type Terminal struct {
	ErrorLog, InfoLog, SuccessLog *log.Logger
	DB                            driver.IDatabase
}
