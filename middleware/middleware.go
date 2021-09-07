package middleware

import "log"

type Middleware struct {
	InfoLog, ErrorLog *log.Logger
}
