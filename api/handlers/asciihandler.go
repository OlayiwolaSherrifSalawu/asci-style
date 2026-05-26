package handlers

import "github.com/OlayiwolaSherrifSalawu/asci-style.git/internal/core"

type AsciHandlers struct {
	Asciservice core.AsciiServiceInterface
}

func NewAsciHandler(ascii core.AsciiServiceInterface) *AsciHandlers {
	return &AsciHandlers{
		Asciservice: ascii,
	}
}
