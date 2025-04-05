package public

import (
	"database/sql"
	"html/template"
	"net/http"
	// "os"
	"strconv"

	"github.com/jmoiron/sqlx"

	"mypersonalwebsite/common"
	"mypersonalwebsite/model"
)

// Path: "/cheatsheets", "/nerdstuff"
//
// this uses entryType to create a list of all entries of that type
// it also uses entryType+"-list" to display content above the list
func RouteListPage(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	locals := common.GetLocals(r, db)
	var q string
	if locals.LoggedIn {
		q = "SELECT * FROM entries WHERE type = ? ORDER BY id DESC"
	} else {
		q = "SELECT * FROM entries WHERE type = ? AND private = 0 ORDER BY id DESC"
	}
	var entries []model.Entry
	err := db.Select(&entries, q, entryType)
	if err != nil && err != sql.ErrNoRows {
		common.SqlError(w, err)
		return
	}

	// get tags for entries
	for i, entry := range entries {
		tags, err := model.GetTagsByEntryId(db, strconv.Itoa(entry.Id))
		if err != nil {
			common.SqlError(w, err)
			return
		}
		entries[i].Tags = tags
	}

	listPageType := entryType + "-list"
	listPage, err := model.GetEntryByType(db, locals, listPageType)

	ref := r.URL.Path

	err = templates["entries-list"].ExecuteTemplate(w, "base", model.ListPageData{
		BasePageData: model.BasePageData{
			Title:  listPage.Title,
			Locals: locals,
		},
		Id:      listPage.Id,
		Content: template.HTML(common.Md2Html(listPage.Content)),
		Ref:     ref,
		Type:    entryType,
		Entries: entries,
	})
	if err != nil {
		common.ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
}
