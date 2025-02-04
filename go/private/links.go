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

func RouteEditLink(
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
		link  model.Link
		title string
		err   error
	)

	if vars["id"] == "new" {
		link = model.Link{
			Id:         0,
			Title:      "",
			Comment:    "",
			ModifiedAt: "",
		}
		title = "New Entry"
	} else {
		link, err = model.GetLinkById(db, vars["id"])
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "404 Not found", http.StatusNotFound)
				fmt.Println(err)
			} else {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				fmt.Println(err)
			}
			return
		}
		title = "Edit Entry"
	}

	allCategories, err := model.GetAllLinkCategories(db)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "404 Not found", http.StatusNotFound)
			fmt.Println(err)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Println(err)
		}
		return
	}

	// preprocesse date
	modifiedAt, _ := time.Parse(time.RFC3339, link.ModifiedAt)

	var content bytes.Buffer
	err = templates["edit-link"].Execute(&content, model.EditLinkPageData{
		Link:          link,
		ModifiedAt:    modifiedAt.Format("2006-01-02 15:04h"),
		Title:         title,
		AllCategories: allCategories,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	public.RenderBaseTemplate(w, templates, title, &content, locals)
}

func RouteUpdateLink(
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
	url := r.FormValue("url")
	title := r.FormValue("title")
	comment := r.FormValue("comment")
	category_id := r.FormValue("categories")
	fmt.Println(url)

	// Validate the form data
	if title == "" || url == "" {
		http.Error(w, "Title and url are required", http.StatusBadRequest)
		return
	}

	timestamp := time.Now().Format(time.RFC3339)
	// var dbId string
	if id == "0" {
		// Insert a new entry into the database
		// var res sql.Result
		_, err = db.Exec(`
			INSERT INTO links (url, title, comment, category_id, created_at, modified_at, user_id)
				VALUES ($1, $2, $3, $4, $5, $5, $6)
		`, url, title, comment, category_id, timestamp, 1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		// InsertId, _ := res.LastInsertId()
		// dbId = fmt.Sprintf("%v", InsertId)
	} else {
		// Update the entry in the database
		_, err = db.Exec(`
			UPDATE links
			SET url = $1, title = $2, comment = $3, category_id = $4, modified_at = $5
			WHERE id = $6
		`, url, title, comment, category_id, timestamp, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		// dbId = id
	}

	http.Redirect(w, r, "/links", 302)
}

func RouteDeleteLink(
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
	// delete link to tag links first
	_, err = db.Exec(`
		DELETE FROM link_to_tag
		WHERE link_id = $1
	`, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	_, err = db.Exec(`
		DELETE FROM links
		WHERE id = $1
	`, id)
	// fmt.Println(res)
	// fmt.Println(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	http.Redirect(w, r, "/links", 302)
}
