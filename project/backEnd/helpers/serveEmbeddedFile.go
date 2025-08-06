package helpers

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func ServeEmbeddedFile(w http.ResponseWriter, r *http.Request, path string) {
	data, err := os.ReadFile("project/" + path)
	if err != nil {
		log.Printf("500 internal error: %v", err)
		ServeError(w, r, http.StatusInternalServerError)
		return
	}

	// Set content-type based on file extension (basic example)
	switch {
	case strings.HasSuffix(path, ".html"):
		w.Header().Set("Content-Type", "text/html")
	case strings.HasSuffix(path, ".css"):
		w.Header().Set("Content-Type", "text/css")
	}

	w.Write(data)
}
