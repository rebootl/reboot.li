package private

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"mypersonalwebsite/common"
	"mypersonalwebsite/model"
)

// Path: "/edit-link-categories"
func RouteEditLinkCategories(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	locals := common.GetLocals(r, db)
	if !locals.LoggedIn {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	categories, err := model.GetAllLinkCategories(db)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Println(err)
		return
	}

	var content bytes.Buffer
	err = templates["edit-link-categories"].Execute(&content, struct {
		LinkCategories []model.LinkCategory
	}{
		LinkCategories: categories,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	common.RenderBaseTemplate(w, templates, "Edit Tags", &content, locals)
}

// Path: "/edit-link-category/{id}"
func RouteEditLinkCategory(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	locals := common.GetLocals(r, db)
	if !locals.LoggedIn {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)

	var linkCategory model.LinkCategory
	var title string
	if vars["id"] == "new" {
		linkCategory = model.LinkCategory{}
		title = "New link category"
	} else {
		var err error
		linkCategory, err = model.GetLinkCategoryById(db, vars["id"])
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			fmt.Println(err)
			return
		}
		title = "Edit link category"
	}

	var content bytes.Buffer
	err := templates["edit-link-category"].Execute(&content, struct {
		Title        string
		LinkCategory model.LinkCategory
	}{
		Title:        title,
		LinkCategory: linkCategory,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	common.RenderBaseTemplate(w, templates, title, &content, locals)
}

// Path: "/update-link-category"
// Method: POST
func RouteUpdateLinkCategory(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	locals := common.GetLocals(r, db)
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

	// Validate the form data
	if name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	if id == "0" {
		_, err = db.Exec(`
			INSERT INTO link_categories (name)
			VALUES ($1)
		`, name)
	} else {
		_, err = db.Exec(`
			UPDATE link_categories
			SET name = $1
			WHERE id = $2
		`, name, id)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	http.Redirect(w, r, "/edit-link-categories", 302)
	return
}

// Path: "/delete-link-category"
// Method: POST
func RouteDeleteLinkCategory(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	locals := common.GetLocals(r, db)
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
		DELETE FROM link_categories
		WHERE id = $1
	`, id)
	fmt.Println(res)
	fmt.Println(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	http.Redirect(w, r, "/edit-link-categories", 302)
}
