package handler

import (
	"net/http"

	"github.com/rmkane/go-wiki-server/internal/model"
	"github.com/rmkane/go-wiki-server/internal/render"
)

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := model.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	render.RenderTemplate(w, "view", p)
}
