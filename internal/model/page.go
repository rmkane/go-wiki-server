package model

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/rmkane/go-wiki-server/internal/constants"
)

type Page struct {
	Title string
	Body  []byte
}

type PageData struct {
	Page      *Page
	CSRFToken string
}

type PagePreview struct {
	Title string
	Body  template.HTML // Prevent HTML escaping
}

// LoadPage loads a wiki page and converts Markdown to HTML
func LoadPage(title string) (*Page, error) {
	filename := GetFilePath(title)
	body, err := os.ReadFile(filename)

	if os.IsNotExist(err) {
		return &Page{Title: title, Body: []byte("")}, nil
	}
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

// Save saves the page content
func (p *Page) Save() error {
	filename := GetFilePath(p.Title)
	return os.WriteFile(filename, p.Body, 0600)
}

// GetFilePath returns the full path for a given title
func GetFilePath(title string) string {
	return filepath.Join(constants.DataDir, fmt.Sprintf("%s.md", title))
}
