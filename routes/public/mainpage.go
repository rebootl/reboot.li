package public

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/jmoiron/sqlx"

	"mypersonalwebsite/common"
	"mypersonalwebsite/model"
)

// Path: "/", "/privacypolicy"
//
// this uses the entryType to determine the page content
func RouteMainPage(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	version := r.URL.Query().Get("version")

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
	common.RenderEntry(w, r, db, templates, entry, version, common.GetLocals(r, db))
}
