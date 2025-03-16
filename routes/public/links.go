package public

import (
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

	err = templates["links"].ExecuteTemplate(w, "base", model.LinkPageData{
		BasePageData: model.BasePageData{
			Title:  linksPage.Title,
			Locals: locals,
		},
		Id:             linksPage.Id,
		Content:        template.HTML(common.Md2Html(linksPage.Content)),
		LinkCategories: linkCategories,
	})
	if err != nil {
		common.ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
}
