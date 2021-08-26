package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/redmejia/terminal/models"
)

func (h *Handler) HandelProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project

	data := json.NewDecoder(r.Body)
	err := data.Decode(&project)
	if err != nil {
		h.ErrorLog.Println(err)
	}

	h.InfoLog.Println(r.Method)

	err = h.DB.InsertNewProject(project)
	if err != nil {
		h.ErrorLog.Println(err)
	}
}
