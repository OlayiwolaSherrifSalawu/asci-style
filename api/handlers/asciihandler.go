package handlers

import (
	"net/http"

	"github.com/OlayiwolaSherrifSalawu/asci-style.git/internal/core"
)

type AsciHandlers struct {
	Asciservice core.AsciiServiceInterface
}

func NewAsciHandler(ascii core.AsciiServiceInterface) *AsciHandlers {
	return &AsciHandlers{
		Asciservice: ascii,
	}
}

func (a *AsciHandlers) ServeAscii(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		// remember to render a 404 page

	}
}
