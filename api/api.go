package api

import (
	"log"

	"github.com/OlayiwolaSherrifSalawu/asci-style.git/api/handlers"
	"github.com/OlayiwolaSherrifSalawu/asci-style.git/fonts"
	"github.com/OlayiwolaSherrifSalawu/asci-style.git/internal/core"
)

func Run() {
	cons, err := ConnectDb("config.env")
	if err != nil {
		log.Fatal(err)
	}
	Db, err := OpenDB(cons)
	if err != nil {
		log.Fatal(err)
		return
	}
	UserHandlers := handlers.NewUserHandler(Db)
	asciiService, err := core.NewAsciiService(fonts.Embed, "store")
	if err != nil {
		log.Fatal(err)
	}
	asciiServiceInt := core.NewAsciiInterface(asciiService)
	ascihandler := handlers.NewAsciHandler(asciiServiceInt)

	NewApp := NewApplication(UserHandlers, ascihandler)
	app := NewApplicationInterFace(NewApp)
	app.Run()
}
