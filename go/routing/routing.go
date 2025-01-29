package routing

import (
	"mypersonalwebsite/auth"
	"mypersonalwebsite/public"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type Route struct {
	Path        string
	EntryType   string
	HandlerFunc func(entryType string, w http.ResponseWriter, r *http.Request, db *sqlx.DB, templates map[string]*template.Template)
	Methods     []string
}

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
