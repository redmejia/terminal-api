package middleware

import "net/http"

func (m *Middleware) Loggers(next http.Handler) http.Handler {
	logerMid := func(w http.ResponseWriter, r *http.Request) {
		m.InfoLog.Println(r.Method)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(logerMid)
}
