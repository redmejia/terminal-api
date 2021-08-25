package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) HandelProject(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
