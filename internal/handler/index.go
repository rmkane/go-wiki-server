package handler

import (
	"net/http"
	"path/filepath"

	"github.com/rmkane/go-wiki-server/internal/constants"
	"github.com/rmkane/go-wiki-server/internal/render"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	files, err := filepath.Glob(filepath.Join(constants.DataDir, "*.txt"))
	if err != nil {
		http.Error(w, "Unable to load pages", http.StatusInternalServerError)
		return
	}

	var pages []string
	for _, file := range files {
		title := filepath.Base(file) // Get filename
		title = title[:len(title)-4] // Remove ".txt" extension
		pages = append(pages, title)
	}

	data := struct {
		Pages []string
	}{
		Pages: pages,
	}

	render.RenderTemplate(w, "index", data)
}
