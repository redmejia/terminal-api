package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) HandleComment(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
