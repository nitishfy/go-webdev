package render

import (
	"bytes"
	"github.com/nitishfy/go-webdev/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// get template from the template cache
	t, ok := app.TemplateCache[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// render the fetched template
	_, err = //fmt.Println(files)
		buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// get all files that ends with .page.gohtml
	files, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return cache, err
	}

	for _, file := range files {
		name := filepath.Base(file)
		tmplSet, err := template.New(name).ParseFiles(file)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			tmplSet, err = tmplSet.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = tmplSet
	}

	return cache, nil
}
