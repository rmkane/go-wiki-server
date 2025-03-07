package handler

import (
	"net/http"

	"github.com/rmkane/go-wiki-server/internal/model"
	"github.com/rmkane/go-wiki-server/internal/render"
	"github.com/rmkane/go-wiki-server/internal/security"
)

func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := model.LoadPage(title)
	if err != nil {
		p = &model.Page{Title: title}
	}

	data := model.PageData{
		Page:      p,
		CSRFToken: security.GenerateCSRFToken(), // Add CSRF token separately
	}

	render.RenderTemplate(w, "edit", data)
}
