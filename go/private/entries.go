package private

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"mypersonalwebsite/auth"
	"mypersonalwebsite/model"
	"mypersonalwebsite/public"
)

func RouteEditEntry(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	locals := auth.GetLocals(r, db)
	if !locals.LoggedIn {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
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
			if err == sql.ErrNoRows {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			fmt.Println(err)
			return
		}
		title = "Edit Entry"
	}

	// get all tags from db
	allTags, err := model.GetAllEntryTags(db)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Println(err)
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
	var content bytes.Buffer
	err = templates["edit-entry"].Execute(&content, model.EditPageData{
		Type:       entryType,
		Entry:      entry,
		ModifiedAt: modifiedAt.Format("2006-01-02 15:04h"),
		Title:      title,
		AllTags:    allTagsSelected,
		Ref:        ref,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	public.RenderBaseTemplate(w, templates, title, &content, locals)
}

func RouteUpdateEntry(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	locals := auth.GetLocals(r, db)
	if !locals.LoggedIn {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")
	private := r.FormValue("private")

	// Validate the form data
	if title == "" || content == "" {
		http.Error(w, "Title and content are required", http.StatusBadRequest)
		return
	}

	privateBool := false
	if private == "on" {
		privateBool = true
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	selectedTagNames := r.Form["tags"]
	err = model.UpdateEntryTags(db, dbId, selectedTagNames)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	ref := r.FormValue("ref")
	http.Redirect(w, r, ref, 302)
}

func RouteDeleteEntry(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	locals := auth.GetLocals(r, db)
	if !locals.LoggedIn {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	id := r.FormValue("id")

	// var res sql.Result
	// delete entry to tag links first
	_, err = db.Exec(`
		DELETE FROM entry_to_tag
		WHERE entry_id = $1
	`, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	_, err = db.Exec(`
		DELETE FROM entries
		WHERE id = $1
	`, id)
	// fmt.Println(res)
	// fmt.Println(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	ref := r.FormValue("ref")
	http.Redirect(w, r, ref, 302)
}
