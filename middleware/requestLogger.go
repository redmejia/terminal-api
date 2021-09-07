package middleware

import "net/http"

func (m *Middleware) RequestLogger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.InfoLog.Println(r.Host, r.Method)
		next.ServeHTTP(w, r)
	})

}
