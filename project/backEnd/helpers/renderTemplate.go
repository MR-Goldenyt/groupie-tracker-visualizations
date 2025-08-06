package helpers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"groupie-tracker-visualizations/project"
)

// renderTemplate parses and executes an HTML template with provided data.
func RenderTemplate(w http.ResponseWriter, r *http.Request, tplPath string, data interface{}) {
	if _, err1 := os.ReadFile("project/" + tplPath); err1 != nil {
		log.Printf("500 internal error: %v", err1)
		ServeError(w, r, http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFS(project.StaticFS, tplPath)
	if err != nil {
		fmt.Println("Template parse error:", err)
		ServeError(w, r, http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		fmt.Println("Template execute error:", err)
	}
}
