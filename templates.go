package main

import "html/template"

var templatePaths = map[string][]string{
	"entry":                {"templates/entry.html", "templates/base.html"},
	"links":                {"templates/links.html", "templates/base.html"},
	"login":                {"templates/login.html", "templates/base.html"},
	"edit-entry":           {"templates/edit-entry.html", "templates/base.html"},
	"edit-tags":            {"templates/edit-tags.html", "templates/base.html"},
	"edit-tag":             {"templates/edit-tag.html", "templates/base.html"},
	"edit-link-categories": {"templates/edit-link-categories.html", "templates/base.html"},
	"edit-link-category":   {"templates/edit-link-category.html", "templates/base.html"},
	"edit-link":            {"templates/edit-link.html", "templates/base.html"},
	"entries-list":         {"templates/entries-list.html", "templates/base.html"},
}

func getTemplate(files []string) *template.Template {
	return template.Must(template.ParseFiles(files...))
}

func loadTemplates() map[string]*template.Template {
	templates := make(map[string]*template.Template)
	for name, paths := range templatePaths {
		templates[name] = getTemplate(paths)
	}
	return templates
}
