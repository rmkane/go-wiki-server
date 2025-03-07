package handler

import (
	"net/http"

	"github.com/rmkane/go-wiki-server/internal/model"
	"github.com/rmkane/go-wiki-server/internal/render"
	"github.com/rmkane/go-wiki-server/internal/security"
)

func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := model.LoadPage(title)
	if err != nil {
		page = &model.Page{Title: title}
	}

	data, err := security.WrapPage(page)
	if err != nil {
		http.Error(w, "Failed to add CSRF token", http.StatusInternalServerError)
	}

	render.RenderTemplate(w, "edit", data)
}
