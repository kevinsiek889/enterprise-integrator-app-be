package handlers

import "fmt"

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Test() {
	fmt.Println("Auth Handler OK")
}