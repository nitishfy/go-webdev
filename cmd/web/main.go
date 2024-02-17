package main

import (
	"github.com/nitishfy/go-webdev/pkg/config"
	"github.com/nitishfy/go-webdev/pkg/handler"
	"github.com/nitishfy/go-webdev/pkg/render"
	"log"
	"net/http"
)

const portNum = ":8080"

func main() {
	app := config.AppConfig{}
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalf("unable to create the template: %v", err)
		return
	}

	app.TemplateCache = tc
	render.NewTemplate(&app)

	http.HandleFunc("/home", handler.Home)
	http.HandleFunc("/about", handler.About)
	http.ListenAndServe(portNum, nil)
}
