package main

import (
	h "github.com/kradnoel/cambiomz/internal/http"
)

func main() {
	server := h.New()
	server.Default()
	server.Run()
}
