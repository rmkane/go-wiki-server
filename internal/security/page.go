package security

import (
	"errors"
	
	"github.com/rmkane/go-wiki-server/internal/model"
)

func WrapPage(page *model.Page) (*model.PageData, error) {
	if page == nil {
		return nil, errors.New("page is nil")
	}
	wrapped := model.PageData{
		Page:      page,
		CSRFToken: GenerateCSRFToken(),
	}
	return &wrapped, nil
}
