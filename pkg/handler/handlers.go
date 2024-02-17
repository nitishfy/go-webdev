package handler

import (
	"github.com/nitishfy/go-webdev/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml")
}

func About(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Hello About")
	render.RenderTemplate(w, "about.page.gohtml")
}
