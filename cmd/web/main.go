package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Define an application struct to hold the application-wide dependencies for the
// web application.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialze a new instance of our application struct, containing the dependencies.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// http.server-struct contains configuration settings
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
