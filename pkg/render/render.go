package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	parsedTemplate, err := template.ParseFiles("./template/"+tmpl, "./template/base.layout.gohtml")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := parsedTemplate.Execute(w, nil); err != nil {
		fmt.Println(err)
		return
	}
}
