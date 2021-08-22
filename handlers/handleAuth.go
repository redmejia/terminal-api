package handlers

import (
	"net/http"
)

func (h *Handler) HandleAuth(w http.ResponseWriter, r *http.Request) {
	h.DB.Save()
	h.SuccessLog.Println("OK")
}
