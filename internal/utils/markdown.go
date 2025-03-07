package utils

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/yuin/goldmark"

	"github.com/rmkane/go-wiki-server/internal/model"
)

func ToMarkdown(page *model.Page) (*model.PagePreview, error) {
	if page == nil {
		return nil, fmt.Errorf("page is nil")
	}
	var buf bytes.Buffer
	if err := goldmark.Convert(page.Body, &buf); err != nil {
		return nil, err
	}
	previewData := model.PagePreview{
		Title: page.Title,
		Body:  template.HTML(buf.Bytes()),
	}
	return &previewData, nil
}
