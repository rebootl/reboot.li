package public

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"mypersonalwebsite/auth"
	"mypersonalwebsite/model"
)

func RenderMainPage(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	baseTemplate *template.Template,
	entryTemplate *template.Template,
) {
	// get the main page content from sqlite database
	var entry model.Entry
	err := db.Get(&entry,
		"SELECT * FROM entries WHERE type = ? AND private = 0 ORDER BY modified_at DESC LIMIT 1",
		entryType,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
			// TODO: return a 404 page
		} else {
			fmt.Println(err)
		}
		return
	}
	renderEntry(w, r, entryTemplate, baseTemplate, entry)
}

func RenderLinksPage(
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	baseTemplate *template.Template,
	linksTemplate *template.Template,
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
	linksTemplate.Execute(&content, linkCategories)

	baseTemplate.Execute(w, template.HTML(content.String()))
}

func RenderListPage(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	baseTemplate *template.Template,
	listPageTemplate *template.Template,
) {
	var entries []model.Entry
	err := db.Select(&entries, "SELECT * FROM entries WHERE type = ? AND private = 0 ORDER BY id DESC", entryType)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
			// TODO: return a 404 page
		} else {
			fmt.Println(err)
		}
		return
	}

	// select all tags from the database
	var tags []struct {
		Id       int    `db:"id"`
		EntryId  int    `db:"entry_id"`
		TagId    int    `db:"id"`
		TagName  string `db:"name"`
		TagColor string `db:"color"`
	}
	err = db.Select(&tags, `
		SELECT et.entry_id, t.id, t.name, t.color
		FROM entry_tags t
		JOIN entry_to_tag et ON t.id = et.tag_id
		WHERE et.entry_id IN (
			SELECT id
			FROM entries
			WHERE type = ? AND private = 0
		)
	`, entryType)
	if err != nil {
		fmt.Println(err)
		return
	}

	// add tags to entries
	for _, tag := range tags {
		for i, entry := range entries {
			if entry.Id == tag.EntryId {
				t := model.Tag{
					Id:    tag.TagId,
					Name:  tag.TagName,
					Color: tag.TagColor,
				}

				entries[i].Tags = append(entries[i].Tags, t)
			}
		}
	}

	var motd bytes.Buffer
	if entryType == "nerdstuff" {
		motdTemplate := template.Must(template.ParseFiles("templates/motd.txt"))
		motdTemplate.Execute(&motd, nil)
	}

	locals := auth.GetLocals(r, db)

	var content bytes.Buffer
	listPageTemplate.Execute(&content, model.ListPageData{
		Entries: entries,
		Motd:    motd.String(),
		Locals:  locals,
	})
	baseTemplate.Execute(w, template.HTML(content.String()))
}

func RenderListEntry(
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	baseTemplate *template.Template,
	entryTemplate *template.Template,
) {
	vars := mux.Vars(r)
	var entry model.Entry
	err := db.Get(&entry, "SELECT * FROM entries WHERE id = ? AND type = 'cheatsheet' AND private = 0", vars["id"])
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
			// TODO: return a 404 page
		} else {
			fmt.Println(err)
		}
		return
	}

	renderEntry(w, r, entryTemplate, baseTemplate, entry)
}

func renderEntry(
	w http.ResponseWriter,
	r *http.Request,
	entryTemplate *template.Template,
	baseTemplate *template.Template,
	entry model.Entry,
) {
	// convert content to html
	// WARNING: apparently markdown does not sanitize the content,
	//          so if we insert content from a random source this is a security risk,
	//          however I'm only planning to insert my own content here for now,
	//          so I leave it like this for now
	htmlContent := markdown.ToHTML([]byte(entry.Content), nil, nil)

	// preprocesse date
	modifiedAt, _ := time.Parse(time.RFC3339, entry.ModifiedAt)

	var content bytes.Buffer
	entryTemplate.Execute(&content, model.EntryPageData{
		Title:      entry.Title,
		Content:    template.HTML(htmlContent),
		ModifiedAt: modifiedAt.Format("2006-01-02 15:04h"),
	})

	baseTemplate.Execute(w, template.HTML(content.String()))
}

func RenderLogin(
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	baseTemplate *template.Template,
	loginTemplate *template.Template,
) {
	locals := auth.GetLocals(r, db)
	var content bytes.Buffer
	loginTemplate.Execute(&content, locals)
	baseTemplate.Execute(w, template.HTML(content.String()))
}
