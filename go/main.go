package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"mypersonalwebsite/auth"
	"mypersonalwebsite/public"
)

func main() {
	r := mux.NewRouter()

	db, err := sqlx.Connect("sqlite3", "db/db.sqlite")
	if err != nil {
		log.Fatalln(err)
	}

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	templates := loadTemplates()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		public.RenderMainPage("maincontent", w, r, db, templates)
	})

	r.HandleFunc("/privacypolicy", func(w http.ResponseWriter, r *http.Request) {
		public.RenderMainPage("privacypolicy", w, r, db, templates)
	})

	r.HandleFunc("/links", func(w http.ResponseWriter, r *http.Request) {
		public.RenderLinksPage(w, r, db, templates)
	})

	r.HandleFunc("/cheatsheets", func(w http.ResponseWriter, r *http.Request) {
		public.RenderListPage("cheatsheet", w, r, db, templates)
	})

	r.HandleFunc("/cheatsheets/{id}", func(w http.ResponseWriter, r *http.Request) {
		public.RenderListEntry(w, r, db, templates)
	})

	r.HandleFunc("/nerdstuff", func(w http.ResponseWriter, r *http.Request) {
		public.RenderListPage("nerdstuff", w, r, db, templates)
	})

	r.HandleFunc("/nerdstuff/{id}", func(w http.ResponseWriter, r *http.Request) {
		public.RenderListEntry(w, r, db, templates)
	})

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		public.RenderLogin(w, r, db, templates)
	}).Methods("GET")

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		auth.CheckLogin(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		auth.Logout(w, r, db)
	})

	r.HandleFunc("/edit-entry", func(w http.ResponseWriter, r *http.Request) {
		// auth.Logout(w, r, db)

	})

	log.Fatal(http.ListenAndServe(":8080", r))
}

/*
	 func loadTemplates() map[string]*template.Template {
		templates := make(map[string]*template.Template)
		templates["base"] = template.Must(template.ParseFiles("templates/index.html"))
		templates["entry"] = template.Must(template.ParseFiles("templates/entry.html"))
		templates["links"] = template.Must(template.ParseFiles("templates/links.html"))
		templates["cheatsheets"] = template.Must(template.ParseFiles("templates/cheatsheets.html"))
		templates["nerdstuff"] = template.Must(template.ParseFiles("templates/nerdstuff.html"))
		templates["login"] = template.Must(template.ParseFiles("templates/login.html"))
		templates["edit-entry"] = template.Must(template.ParseFiles("templates/edit-entry.html"))
		return templates
	}
*/

var templatePaths = map[string]string{
	"base":        "templates/index.html",
	"entry":       "templates/entry.html",
	"links":       "templates/links.html",
	"cheatsheets": "templates/cheatsheets.html",
	"nerdstuff":   "templates/nerdstuff.html",
	"login":       "templates/login.html",
	"edit-entry":  "templates/edit-entry.html",
}

func getTemplate(filePath string) *template.Template {
	return template.Must(template.ParseFiles(filePath))
}

func loadTemplates() map[string]*template.Template {
	templates := make(map[string]*template.Template)
	for name, path := range templatePaths {
		templates[name] = getTemplate(path)
	}
	return templates
}
