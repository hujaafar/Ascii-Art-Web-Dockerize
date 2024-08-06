package functions

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	fmt.Printf("Server running at http://localhost:2004/\n")
	if err := http.ListenAndServe(":2004", nil); err != nil {
		log.Fatal(err)
	}
}
