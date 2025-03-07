package handler

import (
	"net/http"

	"github.com/rmkane/go-wiki-server/internal/model"
	"github.com/rmkane/go-wiki-server/internal/security"
)

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	csrfToken := r.FormValue("csrf_token")
	if !security.ValidateCSRF(csrfToken) {
		http.Error(w, "Invalid CSRF token", http.StatusForbidden)
		return
	}

	body := r.FormValue("body")
	p := &model.Page{Title: title, Body: []byte(body)}
	err = p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
