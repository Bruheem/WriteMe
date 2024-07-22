package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Bruheem/WriteMe/internal/data"
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
		fmt.Fprintf(w, "invalid ID requested")
		return
	}

	book, err := app.models.Books.Get(id)
	if err != nil {
		fmt.Fprintf(w, "error encountered while fetching")
	}

	fmt.Fprintf(w, "Title: %s, Author: %s", book.Title, book.Author)
}

func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Title  string
		Author string
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	b := &data.Book{
		Title:  input.Title,
		Author: input.Author,
	}

	err = app.models.Books.Insert(b)
	if err != nil {
		fmt.Println("Something bad happened:")
		fmt.Println(err)
	}

	fmt.Fprint(w, "Book created successfully")
}

func (app *application) updateBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "updating your book...")
}

func (app *application) deleteBookHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readID(r)
	if err != nil {
		fmt.Fprintf(w, "invalid ID requested")
	}

	err = app.models.Books.Delete(id)

	// needs to be informative
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, "book deleted successfully")
	}
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
