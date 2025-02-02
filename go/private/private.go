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

func RouteEditTags(
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

	allTags, err := model.GetAllEntryTags(db)
	if err != nil {
		// TODO: better error handling, this isn't necessarily a
		// internal server error it could also be just rows not found
		// -> DBReturnErrorHandler
		if err == sql.ErrNoRows {
			http.Error(w, "404 Not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println(err)
		}
		return
	}

	var content bytes.Buffer
	templates["edit-tags"].Execute(&content, struct {
		Tags []model.Tag
	}{
		Tags: allTags,
	})
	public.RenderBaseTemplate(w, templates, "Edit Tags", &content, locals)
}

func RouteEditTag(
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

	var tag model.Tag
	var title string
	if vars["id"] == "new" {
		tag = model.Tag{
			Id:    0,
			Name:  "",
			Color: "",
		}
		title = "New Tag"
	} else {
		var err error
		tag, err = model.GetEntryTagById(db, vars["id"])
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "404 Not found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				fmt.Println(err)
			}
			return
		}
		title = "Edit Tag"
	}

	var content bytes.Buffer
	templates["edit-tag"].Execute(&content, struct {
		Title string
		Tag   model.Tag
	}{
		Title: title,
		Tag:   tag,
	})
	public.RenderBaseTemplate(w, templates, "Edit Tag", &content, locals)
}

func RouteUpdateTag(
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
	name := r.FormValue("name")
	color := r.FormValue("color")

	// Validate the form data
	if name == "" || color == "" {
		http.Error(w, "Name and color are required", http.StatusBadRequest)
		return
	}

	if id == "0" {
		_, err = db.Exec(`
			INSERT INTO entry_tags (name, user_id, color)
			VALUES ($1, $2, $3)
		`, name, 1, color)
	} else {
		_, err = db.Exec(`
			UPDATE entry_tags
			SET name = $1, color = $2
			WHERE id = $3
		`, name, color, id)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	http.Redirect(w, r, "/edit-tags", 302)
	return
}

func RouteDeleteTag(
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

	var res sql.Result
	res, err = db.Exec(`
		DELETE FROM entry_tags
		WHERE id = $1
	`, id)
	fmt.Println(res)
	fmt.Println(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	http.Redirect(w, r, "/edit-tags", 302)
}
