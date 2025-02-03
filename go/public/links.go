package public

import (
	"bytes"
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
	var linkCategories []model.LinkCategory
	err := db.Select(&linkCategories, "SELECT * FROM link_categories ORDER BY name ASC")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(linkCategories)

	// get the links
	for i, category := range linkCategories {
		err := db.Select(&category.Links, "SELECT * FROM links WHERE category_id = ? ORDER BY title ASC", category.Id)
		if err != nil {
			fmt.Println(err)
			return
		}
		linkCategories[i] = category
	}

	var content bytes.Buffer
	templates["links"].Execute(&content, linkCategories)

	RenderBaseTemplate(w, templates, "Links", &content, auth.GetLocals(r, db))
}
