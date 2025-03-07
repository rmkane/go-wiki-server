package handler

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/rmkane/go-wiki-server/internal/constants"
	"github.com/rmkane/go-wiki-server/internal/render"
)

type Pages struct {
	Pages []string
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	files, err := filepath.Glob(filepath.Join(constants.DataDir, "*.md"))
	if err != nil {
		http.Error(w, "Unable to load pages", http.StatusInternalServerError)
		return
	}

	pages := getPageTitles(files)

	log.Println("Pages:", pages)

	data := Pages{
		Pages: pages,
	}

	render.RenderTemplate(w, "index", data)
}

func getPageTitles(files []string) []string {
	var titles = make([]string, len(files))
	for i, file := range files {
		titles[i] = getPageTitle(file)
	}
	return titles
}

func getPageTitle(file string) string {
	basename := filepath.Base(file)
	return strings.TrimSuffix(basename, filepath.Ext(basename))
}
