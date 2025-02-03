package private

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"mypersonalwebsite/auth"
	"mypersonalwebsite/model"
	"mypersonalwebsite/public"
)

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
