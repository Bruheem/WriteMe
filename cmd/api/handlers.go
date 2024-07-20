package main

import (
	"fmt"
	"net/http"
)

// Page handlers
func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	// The home handler is aimed to display the 4 most recent books modified or created
	fmt.Fprintf(w, "Hello this is the home page!")
}

func (app *application) libraryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Displaying your library...")
}

// Book handlers
func (app *application) getBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readID(r)
	if err != nil {
		app.logger.Println(err)
		return
	}

	fmt.Fprintf(w, "Dispalying the %d book", id)
}

func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "creating your book...")
}

func (app *application) updateBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "updating your book...")
}

func (app *application) deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "deleting your book...")
}

// User Authentication handlers
func (app *application) userSignupForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Displaying signing up form")
}

func (app *application) userSignupPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Signing a user...")
}

func (app *application) userLoginForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Displaying login form")
}

func (app *application) userLoginPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logging in the user")
}

func (app *application) userLogoutPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logging out the user")
}
