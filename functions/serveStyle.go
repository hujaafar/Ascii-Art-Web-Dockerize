package functions

import (
    "net/http"
)

func ServeStyle() {
    fs := http.FileServer(http.Dir("./statics"))
    http.Handle("/statics/", http.StripPrefix("/statics/", fs))
}
