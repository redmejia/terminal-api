package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/redmejia/terminal/models"
)

// HandelProject
func (h *Handler) HandelProject(w http.ResponseWriter, r *http.Request) {

	var project models.Project

	switch r.Method {
	case http.MethodGet:
		developerId, ok := r.URL.Query()["devId"]
		if ok {
			// retrive one project
			// http: //127.0.0.1:8080/project?devId=id
			devId, _ := strconv.Atoi(developerId[0])
			developerProjects, err := h.DB.GetProjectsById(int64(devId))
			if err != nil {
				h.ErrorLog.Println(err)
			}

			err = json.NewEncoder(w).Encode(developerProjects)
			if err != nil {
				h.ErrorLog.Println(err)
			}

		} else if r.URL.Path == "/project" {
			// retrive all projects
			// http: //127.0.0.1:8080/project
			projects, err := h.DB.GetProjects()
			if err != nil {
				h.ErrorLog.Println(err)
			}

			err = json.NewEncoder(w).Encode(projects)
			if err != nil {
				h.ErrorLog.Println(err)
			}

		} else {
			return
		}

	case http.MethodPost:
		// Create new project
		// http: //127.0.0.1:8080/project
		data := json.NewDecoder(r.Body)
		err := data.Decode(&project)
		if err != nil {
			h.ErrorLog.Println(err)
		}

		err = h.DB.InsertNewProject(project)
		if err != nil {
			h.ErrorLog.Println(err)
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("new project"))

	case http.MethodPatch:
		// update project
		// http: //127.0.0.1:8080/project
		data := json.NewDecoder(r.Body)
		err := data.Decode(&project)
		if err != nil {
			h.ErrorLog.Println(err)
		}

		err = h.DB.UpdateProject(project)
		if err != nil {
			h.ErrorLog.Println(err)
		}

	case http.MethodDelete:
		// http://127.0.0.1:8080/project?project=id&dev=id
		devId, _ := strconv.Atoi(r.URL.Query().Get("dev"))
		projectId, _ := strconv.Atoi(r.URL.Query().Get("project"))

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
