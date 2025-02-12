package public

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"mypersonalwebsite/auth"
	"mypersonalwebsite/model"
)

func RouteMainPage(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	version := r.URL.Query().Get("version")
	fmt.Println(version)

	// get the main page content from sqlite database
	var entry model.Entry
	err := db.Get(&entry,
		"SELECT * FROM entries WHERE type = ? AND private = 0 ORDER BY modified_at DESC LIMIT 1",
		entryType,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "404 Not found", http.StatusNotFound)
		} else {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		}
		fmt.Println(err)
		return
	}
	renderEntry(w, r, db, templates, entry, version, auth.GetLocals(r, db))
}

func RouteEntry(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	vars := mux.Vars(r)
	locals := auth.GetLocals(r, db)

	version := r.URL.Query().Get("version")
	fmt.Println(version)

	entry, err := model.GetEntryById(db, locals, vars["id"])
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "404 Not found", http.StatusNotFound)
		} else {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		}
		fmt.Println(err)
		return
	}

	renderEntry(w, r, db, templates, entry, version, locals)
}

func renderEntry(
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
	entry model.Entry,
	version string,
	locals model.Locals,
) {
	if version != "" {
		entryVersion, err := model.GetEntryVersion(db, entry.Id, version)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "404 Not found", http.StatusNotFound)
			} else {
				http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			}
			fmt.Println(err)
			return
		}
		entry.Title = entryVersion.Title
		entry.Content = entryVersion.Content
		entry.ModifiedAt = entryVersion.CreatedAt
	}
	// convert content to html
	htmlContent := md2Html(entry.Content)

	// preprocesse date
	modifiedAt, _ := time.Parse(time.RFC3339, entry.ModifiedAt)

	// get version ids
	versionIds, err := model.GetVersionIds(db, entry.Id)

	var content bytes.Buffer
	err = templates["entry"].Execute(&content, model.EntryPageData{
		Id:         entry.Id,
		Title:      entry.Title,
		Content:    template.HTML(htmlContent),
		ModifiedAt: modifiedAt.Format("2006-01-02 15:04h"),
		Tags:       entry.Tags,
		VersionIds: versionIds,
		IsVersion:  version != "",
		Locals:     locals,
	})
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	RenderBaseTemplate(w, templates, entry.Title, &content, locals)
}

func RenderBaseTemplate(
	w http.ResponseWriter,
	templates map[string]*template.Template,
	title string,
	content *bytes.Buffer,
	locals model.Locals,
) {
	err := templates["base"].Execute(w, model.BasePageData{
		Title:   title,
		Content: template.HTML(content.String()),
		Locals:  locals,
	})
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		fmt.Println(err)
	}
}

func md2Html(md string) string {
	// WARNING: apparently markdown does not sanitize the content,
	//          so if we insert content from a random source this is a security risk,
	//          however I'm only planning to insert my own content here for now,
	//          so I leave it like this for now
	return string(markdown.ToHTML([]byte(md), nil, nil))
}
