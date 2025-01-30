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

	var entry model.Entry
	err := db.Get(&entry, "SELECT * FROM entries WHERE id = ?", vars["id"])
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
			// TODO: return a 404 page
		} else {
			fmt.Println(err)
		}
		return
	}

	// convert content to html
	// WARNING: apparently markdown does not sanitize the content,
	//          so if we insert content from a random source this is a security risk,
	//          however I'm only planning to insert my own content here for now,
	//          so I leave it like this for now
	// htmlContent := markdown.ToHTML([]byte(entry.Content), nil, nil)

	// preprocesse date
	modifiedAt, _ := time.Parse(time.RFC3339, entry.ModifiedAt)

	ref := r.URL.Query().Get("ref")
	var content bytes.Buffer
	templates["edit-entry"].Execute(&content, model.EditPageData{
		Id:         entry.Id,
		Title:      entry.Title,
		Content:    entry.Content,
		ModifiedAt: modifiedAt.Format("2006-01-02 15:04h"),
		Ref:        ref,
	})

	templates["base"].Execute(w, template.HTML(content.String()))
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

	ref := r.FormValue("ref")
	http.Redirect(w, r, ref, 302)
}
