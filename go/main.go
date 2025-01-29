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

	var routes = []Route{
		{
			Path:        "/",
			EntryType:   "maincontent",
			HandlerFunc: public.RenderMainPage,
		},
		{
			Path:        "/privacypolicy",
			EntryType:   "privacypolicy",
			HandlerFunc: public.RenderMainPage,
		},
		{
			Path:        "/links",
			EntryType:   "link",
			HandlerFunc: public.RenderLinksPage,
		},
		{
			Path:        "/cheatsheets",
			EntryType:   "cheatsheet",
			HandlerFunc: public.RenderListPage,
		},
		{
			Path:        "/cheatsheets/{id}",
			HandlerFunc: public.RenderListEntry,
		},
		{
			Path:        "/nerdstuff",
			EntryType:   "nerdstuff",
			HandlerFunc: public.RenderListPage,
		},
		{
			Path:        "/nerdstuff/{id}",
			HandlerFunc: public.RenderListEntry,
		},
		{
			Path:        "/login",
			HandlerFunc: public.RenderLogin,
			Methods:     []string{"GET"},
		},
		{
			Path:        "/login",
			HandlerFunc: auth.CheckLogin,
			Methods:     []string{"POST"},
		},
		{
			Path:        "/logout",
			HandlerFunc: auth.Logout,
		},
	}
	loadRoutes(r, routes, db, templates)
	/*
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
	}*/

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
	EntryType   string
	HandlerFunc func(entryType string, w http.ResponseWriter, r *http.Request, db *sqlx.DB, templates map[string]*template.Template)
	Methods     []string
}

// type RenderHandler func(entryType string, w http.ResponseWriter, r *http.Request, db *sqlx.DB, templates map[string]*template.Template)

// func renderWrapper(renderFunc RenderHandler, entryType string, db *sqlx.DB, templates map[string]*template.Template) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		renderFunc(entryType, w, r, db, templates)
// 	}
// }

func loadRoutes(r *mux.Router, routes []Route, db *sqlx.DB, templates map[string]*template.Template) {
	for _, route := range routes {
		handler := func(w http.ResponseWriter, r *http.Request) {
			route.HandlerFunc(route.EntryType, w, r, db, templates)
		}
		if len(route.Methods) > 0 {
			r.HandleFunc(route.Path, handler).Methods(route.Methods...)
		} else {
			r.HandleFunc(route.Path, handler)
		}
	}
}
