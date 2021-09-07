package handlers

import (
	"net/http"
	"strconv"
)

func (h *Handler) HandleLike(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// http://127.0.0.1:8080/project/like?project=id&dev=id
		pid, _ := strconv.Atoi(r.URL.Query().Get("project"))
		did, _ := strconv.Atoi(r.URL.Query().Get("dev"))

		err := h.DB.LikeAProject(int64(pid), int64(did))
		if err != nil {
			h.ErrorLog.Println(err)
		}

	case http.MethodOptions:
		return
	}
}
