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
	renderEntry(w, r, templates, entry, auth.GetLocals(r, db))
}

func RouteListPage(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	locals := auth.GetLocals(r, db)
	var q string
	if locals.LoggedIn {
		q = "SELECT * FROM entries WHERE type = ? ORDER BY id DESC"
	} else {
		q = "SELECT * FROM entries WHERE type = ? AND private = 0 ORDER BY id DESC"
	}
	var entries []model.Entry
	err := db.Select(&entries, q, entryType)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// select all tags from the database
	// HINT: this could be simplified it maybe it is overly optimized
	var tags []struct {
		Id       int    `db:"id"`
		EntryId  int    `db:"entry_id"`
		TagId    int    `db:"id"`
		TagName  string `db:"name"`
		TagColor string `db:"color"`
	}
	err = db.Select(&tags, `
		SELECT et.entry_id, t.id, t.name, t.color
		FROM entry_tags t
		JOIN entry_to_tag et ON t.id = et.tag_id
		WHERE et.entry_id IN (
			SELECT id
			FROM entries
			WHERE type = ?
		)
	`, entryType)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// add tags to entries
	for _, tag := range tags {
		for i, entry := range entries {
			if entry.Id == tag.EntryId {
				t := model.Tag{
					Id:    tag.TagId,
					Name:  tag.TagName,
					Color: tag.TagColor,
				}

				entries[i].Tags = append(entries[i].Tags, t)
			}
		}
	}

	var motd bytes.Buffer
	if entryType == "nerdstuff" {
		motdTemplate := template.Must(template.ParseFiles("templates/motd.txt"))
		motdTemplate.Execute(&motd, nil)
	}

	var entryTypeToTemplateName = map[string]string{
		"nerdstuff":  "nerdstuff",
		"cheatsheet": "cheatsheets",
	}
	templateName := entryTypeToTemplateName[entryType]

	var content bytes.Buffer
	err = templates[templateName].Execute(&content, model.ListPageData{
		Entries: entries,
		Motd:    motd.String(),
		Locals:  locals,
	})
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	var entryTypeToTitle = map[string]string{
		"nerdstuff":  "Nerd stuff",
		"cheatsheet": "Cheat sheets",
	}
	RenderBaseTemplate(w, templates, entryTypeToTitle[entryType], &content, locals)
}

func RouteListEntry(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	vars := mux.Vars(r)
	locals := auth.GetLocals(r, db)
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

	renderEntry(w, r, templates, entry, locals)
}

func renderEntry(
	w http.ResponseWriter,
	r *http.Request,
	templates map[string]*template.Template,
	entry model.Entry,
	locals model.Locals,
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
	err := templates["entry"].Execute(&content, model.EntryPageData{
		Id:         entry.Id,
		Title:      entry.Title,
		Content:    template.HTML(htmlContent),
		ModifiedAt: modifiedAt.Format("2006-01-02 15:04h"),
		Tags:       entry.Tags,
		Locals:     locals,
	})
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	RenderBaseTemplate(w, templates, entry.Title, &content, locals)
}

func RouteLogin(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	locals := auth.GetLocals(r, db)
	var content bytes.Buffer
	err := templates["login"].Execute(&content, locals)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	RenderBaseTemplate(w, templates, "Login", &content, locals)
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
