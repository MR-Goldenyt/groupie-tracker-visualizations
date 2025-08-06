package helpers

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

// serveError reads the error page template and writes it with the given status code.
func ServeError(w http.ResponseWriter, r *http.Request, status int) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	// Load and write the error page content
	data, err := os.ReadFile("project/frontEnd/template/" + strconv.Itoa(status) + ".html")
	if err != nil {

		data1, err1 := os.ReadFile("project/frontEnd/template/500.html")
		if err1 != nil {
			log.Printf("500 internal error: %v", err)
			w.WriteHeader(500)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		log.Printf("500 internal error: %v", err)
		w.WriteHeader(500)
		w.Write(data1)
		return
	}
	w.WriteHeader(status)
	w.Write(data)
}
