package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/redmejia/terminal/models"
)

// HandleComment
func (h *Handler) HandleComment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Make comment
		// http://127.0.0.1:8080/project/comment
		var comment models.Comments

		data := json.NewDecoder(r.Body)
		err := data.Decode(&comment)
		if err != nil {
			h.ErrorLog.Println(err)
		}

		err = h.DB.MakeAComment(comment)
		if err != nil {
			h.ErrorLog.Println(err)
		}

	case http.MethodOptions:
		return

	}
}
