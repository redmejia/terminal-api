package handlers

import "fmt"

func (h *Handler) HandleAuth() {
	fmt.Println("hello")
	h.DB.Save()
}
