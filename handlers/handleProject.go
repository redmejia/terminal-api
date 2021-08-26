package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/redmejia/terminal/models"
)

func (h *Handler) HandelProject(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:

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

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("new project"))

	case http.MethodDelete:
		// http://127.0.0.1:8080/project?project=id&dev=id
		devId, _ := strconv.Atoi(r.URL.Query().Get("dev"))
		projectId, _ := strconv.Atoi(r.URL.Query().Get("project"))

		h.InfoLog.Println(r.Method)

		err := h.DB.DeleteProject(int64(projectId), int64(devId))
		if err != nil {
			h.ErrorLog.Println(err)
		}

		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("deleted"))

	case http.MethodOptions:
		return
	}
}