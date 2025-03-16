package main

import (
	// "fmt"
	// "html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"mypersonalwebsite/config"
)

func main() {
	r := mux.NewRouter()

	db, err := sqlx.Connect("sqlite3", "file:db/db.sqlite?_foreign_keys=on")
	if err != nil {
		log.Fatalln(err)
	}

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	templates := loadTemplates()
	// t, err := template.ParseGlob("./templates/*.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// debug output t
	// fmt.Println(t.DefinedTemplates())

	loadRoutes(r, routes, db, templates)

	port := config.Port
	log.Fatal(http.ListenAndServe(":"+port, r))
}
