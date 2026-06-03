package handlers

import (
	"errors"
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
		a.clientError(w, 404)
		return
	}
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	if text == "" || banner == "" {
		a.clientError(w, 400)
		return
	}
	data := &templateData{}
	data.AsciiHandle.PlainText = text
	data.AsciiHandle.Banner = banner
	result, err := a.Asciservice.BuildAscii(data.AsciiHandle.PlainText, data.AsciiHandle.Banner)
	if err != nil {
		if errors.Is(err, core.Banner_NOT_FOUND) {
			a.serverError(w, err)
			return
		}
		a.clientError(w, 400)
		return
	}
	data.AsciiHandle.AsciiText = result
	// render data to html
}
