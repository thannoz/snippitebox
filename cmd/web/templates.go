package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/thannoz/snippetbox/internal/models"
)

// Define a templateData type to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.
type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

// humanDate returns a nicely formatted date
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// Initialize a template.FuncMap object and store it in a global variable. This is
// essentially a string-keyed map which acts as a lookup between the names of our
// custom template functions and the functions themselves.
var functions = template.FuncMap{
	"humanDate": humanDate,
}

// newTemplateCache automatically parses all templates in pages & partials folder
func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// Use the filepath.Glob() function to get a slice of all filepaths that
	// match the pattern "./ui/html/pages/*.tmpl".
	pages, err := filepath.Glob("./ui/html/pages/*.gohtml")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// Extract the file name (like 'home.tmpl') from the full filepath
		// and assign it to the name variable.
		name := filepath.Base(page)

		// The template.FuncMap must be registered with the template set before you
		// call the ParseFiles() method
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.gohtml")
		if err != nil {
			return nil, err
		}

		// Call ParseGlob() *on this template set* to add any partials
		ts, err = ts.ParseGlob("./ui/html/partials/*.gohtml")
		if err != nil {
			return nil, err
		}

		// Parse the files into a template set.
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add the template set to the map, using the name of the page
		// (like 'home.tmpl') as the key.
		cache[name] = ts
	}
	return cache, nil
}
