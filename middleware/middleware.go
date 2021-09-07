package middleware

import "log"

// Middleware
type Middleware struct {
	InfoLog, ErrorLog *log.Logger
}

// NewMiddleware
func NewMiddleware(infoLog, errLog *log.Logger) *Middleware {
	return &Middleware{
		InfoLog:  infoLog,
		ErrorLog: errLog,
	}
}
