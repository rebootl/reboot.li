package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"mypersonalwebsite/auth"
	// "mypersonalwebsite/config"
	"mypersonalwebsite/model"
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
		renderMainPage("maincontent", w, r, db, baseTemplate, entryTemplate)
	})

	r.HandleFunc("/privacypolicy", func(w http.ResponseWriter, r *http.Request) {
		renderMainPage("privacypolicy", w, r, db, baseTemplate, entryTemplate)
	})

	r.HandleFunc("/links", func(w http.ResponseWriter, r *http.Request) {
		renderLinksPage(w, r, db, baseTemplate, linksTemplate)
	})

	r.HandleFunc("/cheatsheets", func(w http.ResponseWriter, r *http.Request) {
		renderListPage("cheatsheet", w, r, db, baseTemplate, cheatsheetsTemplate)
	})

	r.HandleFunc("/cheatsheets/{id}", func(w http.ResponseWriter, r *http.Request) {
		renderListEntry(w, r, db, baseTemplate, entryTemplate)
	})

	r.HandleFunc("/nerdstuff", func(w http.ResponseWriter, r *http.Request) {
		renderListPage("nerdstuff", w, r, db, baseTemplate, nerdstuffTemplate)
	})

	r.HandleFunc("/nerdstuff/{id}", func(w http.ResponseWriter, r *http.Request) {
		renderListEntry(w, r, db, baseTemplate, entryTemplate)
	})

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		renderLogin(w, r, db, baseTemplate, loginTemplate)
	}).Methods("GET")

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		auth.CheckLogin(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		auth.Logout(w, r, db)
	})
	log.Fatal(http.ListenAndServe(":8080", r))
}

func renderMainPage(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	baseTemplate *template.Template,
	entryTemplate *template.Template,
) {
	// get the main page content from sqlite database
	var entry model.Entry
	err := db.Get(&entry,
		"SELECT * FROM entries WHERE type = ? AND private = 0 ORDER BY modified_at DESC LIMIT 1",
		entryType,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
			// TODO: return a 404 page
		} else {
			fmt.Println(err)
		}
		return
	}
	renderEntry(w, r, entryTemplate, baseTemplate, entry)
}

func renderLinksPage(
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	baseTemplate *template.Template,
	linksTemplate *template.Template,
) {
	// get the link categories from sqlite database
	var linkCategories []model.LinkCategory
	err := db.Select(&linkCategories, "SELECT * FROM link_categories ORDER BY name ASC")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(linkCategories)

	// get the links
	for i, category := range linkCategories {
		err := db.Select(&category.Links, "SELECT * FROM links WHERE category_id = ? ORDER BY title ASC", category.Id)
		if err != nil {
			fmt.Println(err)
			return
		}
		linkCategories[i] = category
	}

	var content bytes.Buffer
	linksTemplate.Execute(&content, linkCategories)

	baseTemplate.Execute(w, template.HTML(content.String()))
}

func renderListPage(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	baseTemplate *template.Template,
	listPageTemplate *template.Template,
) {
	var entries []model.Entry
	err := db.Select(&entries, "SELECT * FROM entries WHERE type = ? AND private = 0 ORDER BY id DESC", entryType)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
			// TODO: return a 404 page
		} else {
			fmt.Println(err)
		}
		return
	}

	var content bytes.Buffer
	listPageTemplate.Execute(&content, entries)
	baseTemplate.Execute(w, template.HTML(content.String()))
}

func renderListEntry(
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	baseTemplate *template.Template,
	entryTemplate *template.Template,
) {
	vars := mux.Vars(r)
	var entry model.Entry
	err := db.Get(&entry, "SELECT * FROM entries WHERE id = ? AND type = 'cheatsheet' AND private = 0", vars["id"])
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
			// TODO: return a 404 page
		} else {
			fmt.Println(err)
		}
		return
	}

	renderEntry(w, r, entryTemplate, baseTemplate, entry)
}

func renderEntry(
	w http.ResponseWriter,
	r *http.Request,
	entryTemplate *template.Template,
	baseTemplate *template.Template,
	entry model.Entry,
) {
	// convert content to html
	// WARNING: apparently markdown does not sanitize the content,
	//          so if we insert content from a random source this is a security risk,
	//          however I'm only planning to insert my own content here for now,
	//          so I leave it like this for now
	htmlContent := markdown.ToHTML([]byte(entry.Content), nil, nil)

	// preprocesse date
	modifiedAt, _ := time.Parse(time.RFC3339, entry.ModifiedAt)

	var content bytes.Buffer
	entryTemplate.Execute(&content, model.EntryPageData{
		Title:      entry.Title,
		Content:    template.HTML(htmlContent),
		ModifiedAt: modifiedAt.Format("2006-01-02 15:04h"),
	})

	baseTemplate.Execute(w, template.HTML(content.String()))
}

func renderLogin(
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	baseTemplate *template.Template,
	loginTemplate *template.Template,
) {
	locals := auth.GetLocals(r, db)
	var content bytes.Buffer
	loginTemplate.Execute(&content, locals)
	baseTemplate.Execute(w, template.HTML(content.String()))
}
