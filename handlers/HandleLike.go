package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) HandleLike(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
