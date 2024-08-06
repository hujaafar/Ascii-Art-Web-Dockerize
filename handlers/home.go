package handlers

import (
	"net/http"
	"text/template"
	"ASCII-GEN/structs"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		renderErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path == "/" {
		tmpl, err := template.ParseFiles("./statics/index.html")
		if err != nil {
			renderErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, structs.PageData{OutputText: ""})
		if err != nil {
			renderErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}

	renderErrorPage(w, "Not Found", http.StatusNotFound)
}
