package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/redmejia/terminal/models"
)

func (h *Handler) HandleComment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:

		var comment models.Comments

		data := json.NewDecoder(r.Body)
		err := data.Decode(&comment)
		if err != nil {
			h.ErrorLog.Println(err)
		}

		h.InfoLog.Println(r.Method)

		err = h.DB.MakeAComment(comment)
		if err != nil {
			h.ErrorLog.Println(err)
		}

	case http.MethodOptions:
		return

	}
}
