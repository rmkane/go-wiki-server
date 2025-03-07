package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templates = findTemplates("templates/*.html")

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, fmt.Sprintf("%s.html", tmpl), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Load templates dynamically
func findTemplates(pattern string) *template.Template {
	paths, err := matchFiles(pattern)
	if err != nil {
		log.Fatal(err)
	}
	return template.Must(template.ParseFiles(paths...))
}

// Get matching file paths
func matchFiles(pattern string) ([]string, error) {
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}
	if files == nil {
		return []string{}, nil
	}
	return files, nil
}
