package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/redmejia/terminal/models"
)

// HandleAuth for regiter and signin
func (h *Handler) HandleAuth(w http.ResponseWriter, r *http.Request) {
	var user models.User

	data := json.NewDecoder(r.Body)
	err := data.Decode(&user)
	if err != nil {
		h.ErrorLog.Println(err)
	}

	h.InfoLog.Println(r.Method)
	err = h.DB.InsertNewDev(user)
	if err != nil {
		h.ErrorLog.Println(err)
	}
}
