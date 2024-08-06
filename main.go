package main

import (
	"ASCII-GEN/functions"
	"ASCII-GEN/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art", handlers.Submit)
	http.HandleFunc("/export", handlers.ExportHandler)
	functions.ServeStyle()
	functions.StartServer()
}
