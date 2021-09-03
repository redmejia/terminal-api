package handlers

import (
	"net/http"
	"strconv"
)

func (h *Handler) HandleLike(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// http://127.0.0.1:8080/project/like?project=id
		id, _ := strconv.Atoi(r.URL.Query().Get("project"))

		h.InfoLog.Println(r.Method)

		h.DB.LikeAProject(int64(id))

		h.InfoLog.Println("...done")

	case http.MethodOptions:
		return
	}
}
