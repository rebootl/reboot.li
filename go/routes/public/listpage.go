package public

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"

	"mypersonalwebsite/auth"
	"mypersonalwebsite/model"
)

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

	// get tags for entries
	for i, entry := range entries {
		tags, err := model.GetTagsByEntryId(db, strconv.Itoa(entry.Id))
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		entries[i].Tags = tags
	}

	var motd bytes.Buffer
	if entryType == "nerdstuff" {
		motdTemplate := template.Must(template.ParseFiles("templates/motd.txt"))
		motdTemplate.Execute(&motd, nil)
	}

	listPageType := entryType + "-list"
	listPage, err := model.GetEntryByType(db, locals, listPageType)

	ref := r.URL.Path

	var content bytes.Buffer
	err = templates["entries-list"].Execute(&content, model.ListPageData{
		Id:      listPage.Id,
		Title:   listPage.Title,
		Motd:    motd.String(),
		Content: template.HTML(md2Html(listPage.Content)),
		Ref:     ref,
		Type:    entryType,
		Entries: entries,
		Locals:  locals,
	})
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	RenderBaseTemplate(w, templates, listPage.Title, &content, locals)
}
