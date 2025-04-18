package private

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"mypersonalwebsite/common"
	"mypersonalwebsite/model"
)

// Path: "/edit-entry/{id}"
func RouteEditEntry(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	locals := common.GetLocals(r, db)
	if !locals.LoggedIn {
		common.ErrorPage(w, nil, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)

	var (
		entry model.Entry
		title string
		err   error
	)

	if vars["id"] == "new" {
		entry = model.Entry{
			Id:         0,
			Title:      "",
			Content:    "",
			Private:    false,
			Tags:       []model.Tag{},
			ModifiedAt: "",
		}
		title = "New Entry"
		entryType = r.URL.Query().Get("type")
	} else {
		entry, err = model.GetEntryById(db, locals, vars["id"])
		if err != nil {
			common.SqlError(w, err)
			return
		}
		title = "Edit Entry"
	}

	// get all tags from db
	allTags, err := model.GetAllEntryTags(db)
	if err != nil {
		common.SqlError(w, err)
		return
	}

	// built structure for multiselect
	var allTagsSelected []model.TagWithStatus
	entryTags := make(map[string]bool)
	for _, tag := range entry.Tags {
		entryTags[tag.Name] = true
	}
	for _, tag := range allTags {
		var tagWithStatus model.TagWithStatus
		if entryTags[tag.Name] {
			tagWithStatus = model.TagWithStatus{
				Tag:      tag,
				Selected: true,
			}
		} else {
			tagWithStatus = model.TagWithStatus{
				Tag:      tag,
				Selected: false,
			}
		}
		allTagsSelected = append(allTagsSelected, tagWithStatus)
	}

	// preprocesse date
	modifiedAt, _ := time.Parse(time.RFC3339, entry.ModifiedAt)

	ref := r.URL.Query().Get("ref")
	err = templates["edit-entry"].ExecuteTemplate(w, "base", model.EditPageData{
		BasePageData: model.BasePageData{
			Title:  title,
			Locals: locals,
		},
		Type:       entryType,
		Entry:      entry,
		ModifiedAt: modifiedAt.Format("2006-01-02 15:04h"),
		AllTags:    allTagsSelected,
		Ref:        ref,
	})
	if err != nil {
		common.ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
}

// Path: "/update-entry"
// Method: POST
func RouteUpdateEntry(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	locals := common.GetLocals(r, db)
	if !locals.LoggedIn {
		common.ErrorPage(w, nil, http.StatusUnauthorized)
		return
	}

	err := r.ParseForm()
	if err != nil {
		common.ErrorPage(w, err, http.StatusBadRequest)
		return
	}
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")
	private := r.FormValue("private")
	version := r.FormValue("version")

	// Validate the form data
	if title == "" || content == "" {
		http.Error(w, "Title and content are required", http.StatusBadRequest)
		return
	}

	privateBool := false
	if private == "on" {
		privateBool = true
	}

	if version == "on" {
		err = model.SaveVersion(db, locals, id)
		if err != nil {
			common.SqlError(w, err)
			return
		}
	}

	timestamp := time.Now().Format(time.RFC3339)
	var dbId string
	if id == "0" {
		entryType = r.FormValue("type")

		// Insert a new entry into the database
		var res sql.Result
		res, err = db.Exec(`
			INSERT INTO entries (title, type, user_id, content, private, created_at, modified_at)
			VALUES ($1, $2, $3, $4, $5, $6, $6)
		`, title, entryType, 1, content, privateBool, timestamp)

		InsertId, _ := res.LastInsertId()
		dbId = fmt.Sprintf("%v", InsertId)
	} else {
		// Update the entry in the database
		_, err = db.Exec(`
			UPDATE entries
			SET title = $1, content = $2, private = $3, modified_at = $4
			WHERE id = $5
		`, title, content, privateBool, timestamp, id)

		dbId = id
	}
	if err != nil {
		common.SqlError(w, err)
		return
	}

	selectedTagNames := r.Form["tags"]
	err = model.UpdateEntryTags(db, dbId, selectedTagNames)
	if err != nil {
		common.SqlError(w, err)
		return
	}

	ref := r.FormValue("ref")
	http.Redirect(w, r, ref, 302)
}

// Path: "/delete-entry"
// Method: POST
func RouteDeleteEntry(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	locals := common.GetLocals(r, db)
	if !locals.LoggedIn {
		common.ErrorPage(w, nil, http.StatusUnauthorized)
		return
	}

	err := r.ParseForm()
	if err != nil {
		common.ErrorPage(w, err, http.StatusBadRequest)
		return
	}
	id := r.FormValue("id")

	// delete entry to tag links first
	_, err = db.Exec(`
		DELETE FROM entry_to_tag
		WHERE entry_id = $1
	`, id)
	if err != nil {
		common.SqlError(w, err)
		return
	}

	var res sql.Result
	res, err = db.Exec(`
		DELETE FROM entries
		WHERE id = $1
	`, id)
	if err != nil {
		common.SqlError(w, err)
		return
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		http.Error(w, "Entry not found", http.StatusNotFound)
		return
	}

	ref := r.FormValue("ref")
	http.Redirect(w, r, ref, 302)
}

// Path: "/delete-version"
// Method: POST
func RouteDeleteVersion(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	locals := common.GetLocals(r, db)
	if !locals.LoggedIn {
		common.ErrorPage(w, nil, http.StatusUnauthorized)
		return
	}

	err := r.ParseForm()
	if err != nil {
		common.ErrorPage(w, err, http.StatusBadRequest)
		return
	}
	id := r.FormValue("id")
	version := r.FormValue("version")

	_, err = db.Exec(`
		DELETE FROM entries_versions
		WHERE entry_id = $1 AND id = $2
	`, id, version)
	if err != nil {
		common.SqlError(w, err)
		return
	}

	ref := r.FormValue("ref")
	http.Redirect(w, r, ref, 302)
}
