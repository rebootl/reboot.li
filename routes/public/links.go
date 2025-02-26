package public

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/jmoiron/sqlx"

	"mypersonalwebsite/common"
	"mypersonalwebsite/model"
)

// Path: "/links"
func RouteLinksPage(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	// get the link categories from sqlite database
	linkCategories, err := model.GetAllLinkCategories(db)
	if err != nil {
		common.SqlError(w, err)
		return
	}

	// get the links
	for i, category := range linkCategories {
		err := db.Select(&category.Links, "SELECT * FROM links WHERE category_id = ? ORDER BY title ASC", category.Id)
		if err != nil {
			common.SqlError(w, err)
			return
		}
		linkCategories[i] = category
	}

	locals := common.GetLocals(r, db)

	linksPage, err := model.GetEntryByType(db, locals, "linkscontent")
	if err != nil {
		common.SqlError(w, err)
		return
	}

	var content bytes.Buffer
	templates["links"].Execute(&content, struct {
		Id             int
		Title          string
		Content        template.HTML
		LinkCategories []model.LinkCategory
		LoggedIn       bool
	}{
		Id:             linksPage.Id,
		Title:          linksPage.Title,
		Content:        template.HTML(common.Md2Html(linksPage.Content)),
		LinkCategories: linkCategories,
		LoggedIn:       locals.LoggedIn,
	})

	common.RenderBaseTemplate(w, templates, linksPage.Title, &content, []string{}, locals)
}
