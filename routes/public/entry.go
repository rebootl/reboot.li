package public

import (
	// "database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"mypersonalwebsite/common"
	"mypersonalwebsite/model"
)

// Path: "/entry/{id}"
func RouteEntry(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	vars := mux.Vars(r)
	locals := common.GetLocals(r, db)

	version := r.URL.Query().Get("version")
	fmt.Println(version)

	entry, err := model.GetEntryById(db, locals, vars["id"])
	if err != nil {
		common.ErrorSQLNotFound(w, err)
		return
	}

	common.RenderEntry(w, r, db, templates, entry, version, locals)
}
