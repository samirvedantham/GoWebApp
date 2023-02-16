package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "This is the home page!")
	RenderTemplate(w, "home.html")

}

// About
func About(w http.ResponseWriter, r *http.Request) {
	//sum := AddValues(2, 2)
	//_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2+2 is: %d", sum))
	RenderTemplate(w, "about.html")
}

/*func Divide(w http.ResponseWriter, r *http.Request) {
	division, err := DivideValues(100.0, 10)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Cannot divide by 0!"))
	}
	fmt.Fprintf(w, fmt.Sprintf("The result is: %f", division))
}*/

// AddValues adds 2 integers and returns the sum
/*func AddValues(x, y int) int {
	return x + y
}*/

/*
func DivideValues(x, y float32) (float32, error) {

	if y <= 0 {
		err := fmt.Errorf("Cannot Divide by 0")
		return 0.0, err
	}
	return x / y, nil
}*/
