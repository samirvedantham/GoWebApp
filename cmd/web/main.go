package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/samirvedantham/GoWebApp/pkg/config"
	"github.com/samirvedantham/GoWebApp/pkg/handlers"
	"github.com/samirvedantham/GoWebApp/pkg/render"
)

const PortNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		n, err := fmt.Fprintf(w, "Hello World!")

		if err != nil {

			fmt.Println(err)
		}

		fmt.Println(fmt.Sprintf("Bytes Written %v", n))

	})*/

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	//http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on Port: %s\n", PortNumber))

	_ = http.ListenAndServe(PortNumber, nil) //if there's an error I don't care

}
