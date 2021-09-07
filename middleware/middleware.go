package middleware

import "log"

type Middleware struct {
	InfoLog, ErrorLog *log.Logger
}

func NewMiddleware(infoLog, errLog *log.Logger) *Middleware {
	return &Middleware{
		InfoLog:  infoLog,
		ErrorLog: errLog,
	}
}
