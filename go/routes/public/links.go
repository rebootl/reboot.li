package public

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/jmoiron/sqlx"

	"mypersonalwebsite/auth"
	"mypersonalwebsite/model"
)

func RouteLinksPage(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	// get the link categories from sqlite database
	linkCategories, err := model.GetAllLinkCategories(db)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// get the links
	for i, category := range linkCategories {
		err := db.Select(&category.Links, "SELECT * FROM links WHERE category_id = ? ORDER BY title ASC", category.Id)
		if err != nil && err != sql.ErrNoRows {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		linkCategories[i] = category
	}

	locals := auth.GetLocals(r, db)

	var content bytes.Buffer
	templates["links"].Execute(&content, struct {
		LinkCategories []model.LinkCategory
		LoggedIn       bool
	}{
		LinkCategories: linkCategories,
		LoggedIn:       locals.LoggedIn,
	})

	RenderBaseTemplate(w, templates, "Links", &content, locals)
}
