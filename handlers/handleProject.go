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

		if r.URL.Query().Has("devID") {
			// retrive  a slice of project by developer id
			// http://127.0.0.1:8080/project?devID=1232
			devID, _ := strconv.Atoi(r.URL.Query().Get("devID"))

			projects, err := h.DB.GetProjectsById(int64(devID))
			if err != nil {
				h.ErrorLog.Println(err)
			}

			err = json.NewEncoder(w).Encode(projects)
			if err != nil {
				h.ErrorLog.Println(err)
			}

		} else if r.URL.Query().Has("proID") {
			// retrive one project
			// http://127.0.0.1:8080/project?proID=id
			proID, _ := strconv.Atoi(r.URL.Query().Get("proID"))

			project, err := h.DB.GetProjectById(int64(proID))
			if err != nil {
				h.ErrorLog.Println(err)
			}

			err = json.NewEncoder(w).Encode(project)
			if err != nil {
				h.ErrorLog.Println(err)
			}

		} else if r.URL.Query().Has("top-projects") {
			// get all top project
			// http://127.0.0.1:8080/project?top-projects=all
			topProjects, err := h.DB.GetTopProject()
			if err != nil {
				h.ErrorLog.Println(err)
			}

			err = json.NewEncoder(w).Encode(topProjects)
			if err != nil {
				h.ErrorLog.Println(err)
			}

		} else if r.URL.Path == "/project" {
			// retrive slice of projects
			// http://127.0.0.1:8080/project
			projects, err := h.DB.GetProjects()
			if err != nil {
				h.ErrorLog.Println(err)
			}

			err = json.NewEncoder(w).Encode(projects)
			if err != nil {
				h.ErrorLog.Println(err)
			}

		} else {
			w.WriteHeader(http.StatusNotFound)
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
