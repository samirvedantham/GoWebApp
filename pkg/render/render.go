package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")

	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template", err)
	}
}

var template_cache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	//var err error

	//check to see if we already have the template in our cache
	_, inMap := template_cache[t]

	if !inMap {
		log.Println("creating template cache")
		err := createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("using cached template")
	}

	tmpl = template_cache[t]

	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {

	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
		"./templates/home.page.tmpl",
	}

	//parse the template

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	template_cache[t] = tmpl

	return nil

}
