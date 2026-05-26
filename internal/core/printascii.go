package core

import (
	"embed"
	"strings"
)

type AsciiServiceInterface interface {
	BuildAscii(file, banner string) (string, error)
}
type AsciiService struct {
	Banners map[string][]string
}

func NewAsciiService(embedss embed.FS, directoryname string) (*AsciiService, error) {
	embedd := make(map[string][]string)
	files, err := embedss.ReadDir(directoryname)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		vals, _ := embedss.ReadFile(file.Name())
		embedd[strings.TrimSuffix(file.Name(), ".")] = strings.Split(string(vals), "\n")
	}
	return &AsciiService{
		Banners: embedd,
	}, nil
}

func NewAsciiInterface(asci *AsciiService) AsciiServiceInterface {
	return asci
}
