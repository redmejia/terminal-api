package handlers

import (
	"net/http"
)

// HandleAuth for regiter and signin
func (h *Handler) HandleAuth(w http.ResponseWriter, r *http.Request) {
	h.DB.Save()
	h.SuccessLog.Println("OK")
}
