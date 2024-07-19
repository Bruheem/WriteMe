package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// The home handler is aimed to display the 4 most recent books modified or created
	fmt.Fprintf(w, "Hello this is the home page!")
}

func (app *application) getBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readID(r)
	if err != nil {
		app.logger.Println(err)
		return
	}

	fmt.Fprintf(w, "Dispalying the %d book", id)
}
