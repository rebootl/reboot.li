package private

import (
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
		common.ErrorPage(w, nil, http.StatusUnauthorized)
		return
	}

	categories, err := model.GetAllLinkCategories(db)
	if err != nil {
		common.SqlError(w, err)
		return
	}

	err = templates["edit-link-categories"].ExecuteTemplate(w, "base", struct {
		model.BasePageData
		LinkCategories []model.LinkCategory
	}{
		BasePageData: model.BasePageData{
			Title:  "Edit Link Categories",
			Locals: locals,
		},
		LinkCategories: categories,
	})
	if err != nil {
		common.ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
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
		common.ErrorPage(w, nil, http.StatusUnauthorized)
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
			common.SqlError(w, err)
			return
		}
		title = "Edit link category"
	}

	err := templates["edit-link-category"].ExecuteTemplate(w, "base", struct {
		model.BasePageData
		LinkCategory model.LinkCategory
	}{
		BasePageData: model.BasePageData{
			Title:  title,
			Locals: locals,
		},
		LinkCategory: linkCategory,
	})
	if err != nil {
		common.ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
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
		common.ErrorPage(w, nil, http.StatusUnauthorized)
		return
	}

	err := r.ParseForm()
	if err != nil {
		common.ErrorPage(w, err, http.StatusBadRequest)
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
		common.SqlError(w, err)
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
		common.ErrorPage(w, nil, http.StatusUnauthorized)
		return
	}

	err := r.ParseForm()
	if err != nil {
		common.ErrorPage(w, err, http.StatusBadRequest)
		return
	}
	id := r.FormValue("id")

	_, err = db.Exec(`
		DELETE FROM link_categories
		WHERE id = $1
	`, id)
	if err != nil {
		common.SqlError(w, err)
		return
	}
	http.Redirect(w, r, "/edit-link-categories", 302)
}
