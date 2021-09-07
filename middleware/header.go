package middleware

import "net/http"

// Header  set json header
func (m *Middleware) Header(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})

}

// JsonFormat check json format
func (m *Middleware) JsonFormat(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Content-Type") != "application/json" {
			m.ErrorLog.Println("header not json")
			return
		}
		next.ServeHTTP(w, r)

	})

}
