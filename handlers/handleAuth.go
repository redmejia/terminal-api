package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/redmejia/terminal/models"
)

// HandleAuth for regiter and signin
func (h *Handler) HandleAuth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// http://127.0.0.1:8080/
		var user models.User

		data := json.NewDecoder(r.Body)
		err := data.Decode(&user)
		if err != nil {
			h.ErrorLog.Println(err)
		}

		err = h.DB.InsertNewDev(user)
		if err != nil {
			h.ErrorLog.Println(err)
		}

	case http.MethodOptions:
		return

	}
}
