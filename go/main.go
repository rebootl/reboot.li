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

	baseTemplate := template.Must(template.ParseFiles("templates/index.html"))
	entryTemplate := template.Must(template.ParseFiles("templates/entry.html"))
	linksTemplate := template.Must(template.ParseFiles("templates/links.html"))
	cheatsheetsTemplate := template.Must(template.ParseFiles("templates/cheatsheets.html"))
	nerdstuffTemplate := template.Must(template.ParseFiles("templates/nerdstuff.html"))
	loginTemplate := template.Must(template.ParseFiles("templates/login.html"))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		public.RenderMainPage("maincontent", w, r, db, baseTemplate, entryTemplate)
	})

	r.HandleFunc("/privacypolicy", func(w http.ResponseWriter, r *http.Request) {
		public.RenderMainPage("privacypolicy", w, r, db, baseTemplate, entryTemplate)
	})

	r.HandleFunc("/links", func(w http.ResponseWriter, r *http.Request) {
		public.RenderLinksPage(w, r, db, baseTemplate, linksTemplate)
	})

	r.HandleFunc("/cheatsheets", func(w http.ResponseWriter, r *http.Request) {
		public.RenderListPage("cheatsheet", w, r, db, baseTemplate, cheatsheetsTemplate)
	})

	r.HandleFunc("/cheatsheets/{id}", func(w http.ResponseWriter, r *http.Request) {
		public.RenderListEntry(w, r, db, baseTemplate, entryTemplate)
	})

	r.HandleFunc("/nerdstuff", func(w http.ResponseWriter, r *http.Request) {
		public.RenderListPage("nerdstuff", w, r, db, baseTemplate, nerdstuffTemplate)
	})

	r.HandleFunc("/nerdstuff/{id}", func(w http.ResponseWriter, r *http.Request) {
		public.RenderListEntry(w, r, db, baseTemplate, entryTemplate)
	})

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		public.RenderLogin(w, r, db, baseTemplate, loginTemplate)
	}).Methods("GET")

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		auth.CheckLogin(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		auth.Logout(w, r, db)
	})
	log.Fatal(http.ListenAndServe(":8080", r))
}
