package handlers

import (
	"net/http"
	"text/template"
	"ASCII-GEN/structs"  
	"ASCII-GEN/functions"  
)

func Submit(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.NotFound(w, r)
		return	
	}
	if r.Method != "POST" {
		renderErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	
	err := r.ParseForm()
	if err != nil {
		renderErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	text := r.FormValue("text")
	format := r.FormValue("format")
	output, err := functions.Ascii(text, format)
	if text == "" || len(text) > 250 || (format != "standard" && format != "shadow" && format != "thinkertoy") {
		renderErrorPage(w, "Bad Request: Check your input", http.StatusBadRequest)
		return
	}

	if err != nil {
		if err.Error() == "error: invalid character" {
			renderErrorPage(w, "Bad Request", http.StatusBadRequest)
		} else {
			renderErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	data := structs.PageData{
		OutputText: output,
	}

	tmpl, err := template.ParseFiles("./statics/index.html")
	if err != nil {
		renderErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		renderErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
