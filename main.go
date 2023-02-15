package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const PortNumber = ":8080"

func main() {

	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		n, err := fmt.Fprintf(w, "Hello World!")

		if err != nil {

			fmt.Println(err)
		}

		fmt.Println(fmt.Sprintf("Bytes Written %v", n))

	})*/

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	_, _ = fmt.Printf(fmt.Sprintf("Starting application on Port: %s", PortNumber))

	_ = http.ListenAndServe(PortNumber, nil) //if there's an error I don't care

}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	err := parsedTemplate.Execute(w, nil)

	if err != nil {

		fmt.Println("error parsing template")
		return
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page!")
	RenderTemplate(w, "home.tmpl")

}

func Divide(w http.ResponseWriter, r *http.Request) {
	division, err := DivideValues(100.0, 10)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Cannot divide by 0!"))
	}
	fmt.Fprintf(w, fmt.Sprintf("The result is: %f", division))
}

// About
func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2+2 is: %d", sum))

}

// AddValues adds 2 integers and returns the sum
func AddValues(x, y int) int {
	return x + y
}

func DivideValues(x, y float32) (float32, error) {

	if y <= 0 {
		err := fmt.Errorf("Cannot Divide by 0")
		return 0.0, err
	}
	return x / y, nil
}
