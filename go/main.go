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

	var publicRoutes = []Route{
		{
			Path: "/",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				public.RenderMainPage("maincontent", w, r, db, templates)
			},
		},
		{
			Path: "/privacypolicy",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				public.RenderMainPage("privacypolicy", w, r, db, templates)
			},
		},
		{
			Path: "/links",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				public.RenderLinksPage(w, r, db, templates)
			},
		},
		{
			Path: "/cheatsheets",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				public.RenderListPage("cheatsheet", w, r, db, templates)
			},
		},
		{
			Path: "/cheatsheets/{id}",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				public.RenderListEntry(w, r, db, templates)
			},
		},
		{
			Path: "/nerdstuff",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				public.RenderListPage("nerdstuff", w, r, db, templates)
			},
		},
		{
			Path: "/nerdstuff/{id}",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				public.RenderListEntry(w, r, db, templates)
			},
		},
		{
			Path: "/login",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				public.RenderLogin(w, r, db, templates)
			},
			Methods: []string{"GET"},
		},
		{
			Path: "/login",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				auth.CheckLogin(w, r, db)
			},
			Methods: []string{"POST"},
		},
		{
			Path: "/logout",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				auth.Logout(w, r, db)
			},
		},
		{
			Path: "/edit-entry",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				// auth.Logout(w, r, db)
			},
		},
	}
	loadRoutes(r, publicRoutes)

	log.Fatal(http.ListenAndServe(":8080", r))
}

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

type Route struct {
	Path        string
	HandlerFunc http.HandlerFunc
	Methods     []string
}

func loadRoutes(r *mux.Router, routes []Route) {
	for _, route := range routes {
		if len(route.Methods) > 0 {
			r.HandleFunc(route.Path, route.HandlerFunc).Methods(route.Methods...)
		} else {
			r.HandleFunc(route.Path, route.HandlerFunc)
		}
	}
}
