package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/redmejia/terminal/models"
)

func (h *Handler) HandelProject(w http.ResponseWriter, r *http.Request) {

	var project models.Project

	switch r.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")

		projectId, ok := r.URL.Query()["project"]
		if ok {
			// http: //127.0.0.1:8080/project?project=id
			projectId, _ := strconv.Atoi(projectId[0])
			singleProject, err := h.DB.GetProjectById(int64(projectId))
			if err != nil {
				h.ErrorLog.Println(err)
			}

			h.InfoLog.Println(r.Method)

			err = json.NewEncoder(w).Encode(singleProject)
			if err != nil {
				h.ErrorLog.Println(err)
			}

		} else if r.URL.Path == "/project" {

			projects, err := h.DB.GetProjects()
			if err != nil {
				h.ErrorLog.Println(err)
			}

			h.InfoLog.Println(r.Method)

			err = json.NewEncoder(w).Encode(projects)
			if err != nil {
				h.ErrorLog.Println(err)
			}

		} else {
			return
		}

	case http.MethodPost:

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
	case http.MethodPatch:

		data := json.NewDecoder(r.Body)
		err := data.Decode(&project)
		if err != nil {
			h.ErrorLog.Println(err)
		}

		h.InfoLog.Println(r.Method)
		err = h.DB.UpdateProject(project)
		if err != nil {
			h.ErrorLog.Println(err)
		}

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
