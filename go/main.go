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
	_ "github.com/mattn/go-sqlite3"
)

type Entry struct {
	Id         int
	UserId     int
	Type       string
	CreatedAt  string
	ModifiedAt string
	Title      string
	Content    string
	Private    bool
}

type EntryPageData struct {
	Title      string
	Content    template.HTML
	ModifiedAt string
}

func main() {
	r := mux.NewRouter()

	// Initialize the SQLite database
	db, err := sql.Open("sqlite3", "db/db.sqlite")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	baseTemplate := template.Must(template.ParseFiles("templates/index.html"))
	entryTemplate := template.Must(template.ParseFiles("templates/entry.html"))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderMainPage("maincontent", w, r, db, baseTemplate, entryTemplate)
	})

	r.HandleFunc("/privacypolicy", func(w http.ResponseWriter, r *http.Request) {
		renderMainPage("privacypolicy", w, r, db, baseTemplate, entryTemplate)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}

func renderMainPage(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sql.DB,
	baseTemplate *template.Template,
	entryTemplate *template.Template,
) {
	// get the main page content from sqlite database
	var entry Entry
	err := db.QueryRow("SELECT * FROM entries WHERE type = ? AND user_id = 1 AND private = 0", entryType).Scan(
		&entry.Id,
		&entry.UserId,
		&entry.Type,
		&entry.CreatedAt,
		&entry.ModifiedAt,
		&entry.Title,
		&entry.Content,
		&entry.Private,
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

	// convert content to html
	// WARNING: apparently markdown does not sanitize the content,
	//          so if we insert content from a random source this is a security risk,
	//          however I'm only planning to insert my own content here for now,
	//          so I leave it like this for now
	htmlContent := markdown.ToHTML([]byte(entry.Content), nil, nil)

	// preprocesse date
	modifiedAt, _ := time.Parse(time.RFC3339, entry.ModifiedAt)

	var content bytes.Buffer
	entryTemplate.Execute(&content, EntryPageData{
		Title:      entry.Title,
		Content:    template.HTML(htmlContent),
		ModifiedAt: modifiedAt.Format("2006-01-02 15:04h"),
	})

	baseTemplate.Execute(w, template.HTML(content.String()))
}
