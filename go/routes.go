package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"mypersonalwebsite/auth"
	"mypersonalwebsite/private"
	"mypersonalwebsite/public"
)

type Route struct {
	Path        string
	EntryType   string
	HandlerFunc func(entryType string, w http.ResponseWriter, r *http.Request, db *sqlx.DB, templates map[string]*template.Template)
	Methods     []string
}

var routes = []Route{
	{
		Path:        "/",
		EntryType:   "maincontent",
		HandlerFunc: public.RouteMainPage,
	},
	{
		Path:        "/privacypolicy",
		EntryType:   "privacypolicy",
		HandlerFunc: public.RouteMainPage,
	},
	{
		Path:        "/links",
		EntryType:   "link",
		HandlerFunc: public.RouteLinksPage,
	},
	{
		Path:        "/cheatsheets",
		EntryType:   "cheatsheet",
		HandlerFunc: public.RouteListPage,
	},
	{
		Path:        "/entry/{id}",
		HandlerFunc: public.RouteListEntry,
	},
	{
		Path:        "/nerdstuff",
		EntryType:   "nerdstuff",
		HandlerFunc: public.RouteListPage,
	},
	// {
	// 	Path:        "/nerdstuff/{id}",
	// 	HandlerFunc: public.RouteListEntry,
	// },
	{
		Path:        "/login",
		HandlerFunc: public.RouteLogin,
		Methods:     []string{"GET"},
	},
	{
		Path:        "/login",
		HandlerFunc: auth.RouteCheckLogin,
		Methods:     []string{"POST"},
	},
	{
		Path:        "/logout",
		HandlerFunc: auth.RouteLogout,
	},
	{
		Path:        "/edit-entry/{id}",
		HandlerFunc: private.RouteEditEntry,
	},
	{
		Path:        "/update-entry",
		HandlerFunc: private.RouteUpdateEntry,
		Methods:     []string{"POST"},
	},
	{
		Path:        "/delete-entry",
		HandlerFunc: private.RouteDeleteEntry,
		Methods:     []string{"POST"},
	},
	{
		Path:        "/edit-tags",
		HandlerFunc: private.RouteEditTags,
	},
	{
		Path:        "/edit-tag/{id}",
		HandlerFunc: private.RouteEditTag,
	},
	{
		Path:        "/update-tag",
		HandlerFunc: private.RouteUpdateTag,
	},
	{
		Path:        "/delete-tag",
		HandlerFunc: private.RouteDeleteTag,
		Methods:     []string{"POST"},
	},
	{
		Path:        "/edit-link-categories",
		HandlerFunc: private.RouteEditLinkCategories,
	},
	{
		Path:        "/edit-link-category/{id}",
		HandlerFunc: private.RouteEditLinkCategory,
	},
	{
		Path:        "/update-link-category",
		HandlerFunc: private.RouteUpdateLinkCategory,
		Methods:     []string{"POST"},
	},
	{
		Path:        "/delete-link-category",
		HandlerFunc: private.RouteDeleteLinkCategory,
		Methods:     []string{"POST"},
	},
	{
		Path:        "/edit-link/{id}",
		HandlerFunc: private.RouteEditLink,
	},
	{
		Path:        "/update-link",
		HandlerFunc: private.RouteUpdateLink,
		Methods:     []string{"POST"},
	},
	{
		Path:        "/delete-link",
		HandlerFunc: private.RouteDeleteLink,
		Methods:     []string{"POST"},
	},
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
