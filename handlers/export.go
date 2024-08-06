package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	"ASCII-GEN/structs"
)

var tmpl = template.Must(template.ParseFiles("./statics/index.html"))

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		renderErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		renderErrorPage(w, fmt.Sprintf("Error parsing form data: %v", err), http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	if text == "" {
		renderErrorPage(w, "Input text cannot be empty", http.StatusBadRequest)
		return
	}

	// Replace this with actual ASCII art generation logic
	output := "Generated ASCII Art for: " + text

	data := structs.PageData{OutputText: output}
	err := tmpl.Execute(w, data)
	if err != nil {
		renderErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func ExportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		renderErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/export" {
		renderErrorPage(w, "Not Found", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		renderErrorPage(w, fmt.Sprintf("Error parsing form data: %v", err), http.StatusBadRequest)
		return
	}

	output := r.FormValue("output")
	if output == "" {
		renderErrorPage(w, "No output provided", http.StatusBadRequest)
		return
	}

	format := r.FormValue("format")
	if format != "txt" {
		renderErrorPage(w, "Unsupported format", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", "attachment; filename=ASCII-GEN.txt")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(output)))

	_, err := w.Write([]byte(output))
	if err != nil {
		renderErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
