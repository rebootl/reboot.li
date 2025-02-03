package main

import "html/template"

var templatePaths = map[string]string{
	"base":                 "templates/index.html",
	"entry":                "templates/entry.html",
	"links":                "templates/links.html",
	"cheatsheets":          "templates/cheatsheets.html",
	"nerdstuff":            "templates/nerdstuff.html",
	"login":                "templates/login.html",
	"edit-entry":           "templates/edit-entry.html",
	"edit-tags":            "templates/edit-tags.html",
	"edit-tag":             "templates/edit-tag.html",
	"edit-link-categories": "templates/edit-link-categories.html",
	"edit-link-category":   "templates/edit-link-category.html",
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
