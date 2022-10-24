package main

import (
	"net/http"

	"github.com/justinas/alice"
)

// The routes() method returns a servermux containing out application routes.
func (app *application) routes() http.Handler {

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// removed recoverPanic-middleware
	standard := alice.New(app.logRequest, secureHeaders)

	return standard.Then(mux)
}
