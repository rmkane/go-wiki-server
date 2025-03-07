package handler

import (
	"net/http"

	"github.com/rmkane/go-wiki-server/internal/model"
	"github.com/rmkane/go-wiki-server/internal/render"
	"github.com/rmkane/go-wiki-server/internal/utils"
)

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := model.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	data, err := utils.ToMarkdown(page)
	render.RenderTemplate(w, "view", data)
}
