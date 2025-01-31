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

	entry, err := model.GetEntryById(db, locals, vars["id"])
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
			http.Error(w, "404 Not found", http.StatusNotFound)
		} else {
			fmt.Println(err)
		}
		return
	}

	// get all tags from db
	var allTags []model.Tag
	err = db.Select(&allTags, "SELECT * FROM entry_tags")
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
	templates["edit-entry"].Execute(&content, model.EditPageData{
		Id:         entry.Id,
		Title:      entry.Title,
		Content:    entry.Content,
		Private:    entry.Private,
		ModifiedAt: modifiedAt.Format("2006-01-02 15:04h"),
		Tags:       entry.Tags,
		AllTags:    allTagsSelected,
		Ref:        ref,
	})

	public.RenderBaseTemplate(w, templates, "Edit Entry", &content, locals)
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
	// Update the entry in the database
	_, err = db.Exec(`
		UPDATE entries
		SET title = $1, content = $2, private = $3, modified_at = $4
		WHERE id = $5
	`, title, content, privateBool, timestamp, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	selectedTagNames := r.Form["tags"]
	err = model.UpdateEntryTags(db, id, selectedTagNames)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	ref := r.FormValue("ref")
	http.Redirect(w, r, ref, 302)
}
