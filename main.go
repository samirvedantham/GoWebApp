package main

import (
	"fmt"
	"net/http"
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
