package private

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"mypersonalwebsite/common"
	"mypersonalwebsite/model"
)

// Path: "/edit-tags"
func RouteEditTags(
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

	allTags, err := model.GetAllEntryTags(db)
	if err != nil {
		common.SqlError(w, err)
		return
	}

	err = templates["edit-tags"].ExecuteTemplate(w, "base", struct {
		model.BasePageData
		Tags []model.Tag
	}{
		BasePageData: model.BasePageData{
			Title:  "Edit Tags",
			Locals: locals,
		},
		Tags: allTags,
	})
	if err != nil {
		common.ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
}

// Path: "/edit-tag/{id}"
func RouteEditTag(
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
			common.SqlError(w, err)
			return
		}
		title = "Edit Tag"
	}

	err := templates["edit-tag"].ExecuteTemplate(w, "base", struct {
		model.BasePageData
		Tag model.Tag
	}{
		BasePageData: model.BasePageData{
			Title:  title,
			Locals: locals,
		},
		Tag: tag,
	})
	if err != nil {
		common.ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
}

// Path: "/update-tag"
// Method: POST
func RouteUpdateTag(
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
		common.SqlError(w, err)
		return
	}
	http.Redirect(w, r, "/edit-tags", 302)
	return
}

// Path: "/delete-tag"
// Method: POST
func RouteDeleteTag(
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
		DELETE FROM entry_tags
		WHERE id = $1
	`, id)
	if err != nil {
		common.SqlError(w, err)
		return
	}
	http.Redirect(w, r, "/edit-tags", 302)
}
