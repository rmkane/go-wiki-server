package model

import (
	"fmt"
	"os"
	"path/filepath"
)

const DataDir = "data"

type PageData struct {
	Page      *Page
	CSRFToken string
}

type Page struct {
	Title     string
	Body      []byte
	CSRFToken string
}

// LoadPage loads a wiki page from a file
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
	return filepath.Join(DataDir, fmt.Sprintf("%s.txt", title))
}
