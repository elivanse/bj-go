package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var templates *template.Template

func indexPage(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "ppal", nil)
}

func main() {
	templates = template.Must(templates.ParseGlob("html/*.html"))

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", indexPage).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./html/")))

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
