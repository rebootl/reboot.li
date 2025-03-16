package common

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"mypersonalwebsite/config"
	"mypersonalwebsite/model"
)

func getMessage(msg string, err error) string {
	if config.Mode == config.ModeDev && err != nil {
		return msg + ": " + err.Error()
	}
	return msg
}

func SqlError(w http.ResponseWriter, err error) {
	if err == sql.ErrNoRows {
		ErrorPage(w, err, http.StatusNotFound)
	} else {
		ErrorPage(w, err, http.StatusInternalServerError)
	}
}

func ErrorPage(w http.ResponseWriter, err error, httpError int) {
	var msg string
	switch httpError {
	case http.StatusNotFound:
		msg = "404 Not Found"
	case http.StatusUnauthorized:
		msg = "401 Unauthorized"
	default:
		msg = "500 Internal Server Error"
	}
	http.Error(w, getMessage(msg, err), httpError)
	fmt.Println(err)
}

func GetLocals(r *http.Request, db *sqlx.DB) model.Locals {
	// Check if the user is logged in
	// NOTE: if it can't find a cookie this will return an error
	cookie, err := r.Cookie(config.CookieName)
	if err != nil {
		// fmt.Println("No cookie found")
		return model.Locals{LoggedIn: false, UserName: ""}
	}

	// Get the session from the database
	// NOTE: if no rows are found it will also return an error
	var session model.Session
	err = db.Get(&session, "SELECT * FROM sessions WHERE uuid = ?", cookie.Value)
	if err != nil {
		// fmt.Println("No session found")
		return model.Locals{LoggedIn: false, UserName: ""}
	}

	// Get the user from the database
	// NOTE: if no rows are found it will also return an error
	var user model.User
	err = db.Get(&user, "SELECT * FROM users WHERE id = ?", session.UserId)
	if err != nil {
		// fmt.Println("No user found")
		return model.Locals{LoggedIn: false, UserName: ""}
	}

	return model.Locals{LoggedIn: true, UserName: user.UserName}
}

func RenderEntry(
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
	entry model.Entry,
	version string,
	locals model.Locals,
) {
	if version != "" {
		entryVersion, err := model.GetEntryVersion(db, entry.Id, version)
		if err != nil {
			SqlError(w, err)
			return
		}
		entry.Title = entryVersion.Title
		entry.Content = entryVersion.Content
		entry.ModifiedAt = entryVersion.LastModifiedAt
	}
	// convert content to html
	htmlContent := Md2Html(entry.Content)

	// preprocesse date
	modifiedAt, _ := time.Parse(time.RFC3339, entry.ModifiedAt)

	// get version ids
	versions, err := getVersions(db, entry.Id, version)
	if err != nil {
		ErrorPage(w, err, http.StatusInternalServerError)
		return
	}

	err = templates["entry"].ExecuteTemplate(w, "base", model.EntryPageData{
		BasePageData: model.BasePageData{
			Title:  entry.Title,
			Locals: locals,
		},
		Content:    template.HTML(htmlContent),
		Id:         entry.Id,
		ModifiedAt: modifiedAt.Format("2006-01-02 15:04h"),
		Tags:       entry.Tags,
		IsVersion:  version != "",
		Versions:   versions,
	})
	if err != nil {
		ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
}

/*
func RenderBaseTemplate(
	w http.ResponseWriter,
	templates map[string]*template.Template,
	title string,
	content *bytes.Buffer,
	locals model.Locals,
) {
	err := templates["base"].Execute(w, model.BasePageData{
		Title: title,
		// Content: template.HTML(content.String()),
		Locals: locals,
	})
	if err != nil {
		ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
}
*/

func Md2Html(md string) string {
	// WARNING: apparently markdown does not sanitize the content,
	//          so if we insert content from a random source this is a security risk,
	//          however I'm only planning to insert my own content here for now,
	//          so I leave it like this for now
	return string(markdown.ToHTML([]byte(md), nil, nil))
}

// getVersions returns the version ids of an entry and the previous and next version of a given version
// this is used to create the "previous" and "next" links in the entry page
func getVersions(db *sqlx.DB, entryId int, version string) (model.PageVersions, error) {
	var v model.PageVersions
	versionIds, err := model.GetVersionIds(db, entryId)
	if err != nil {
		return v, err
	}
	v.VersionIds = versionIds
	// zeroing is not needed because the struct is already zeroed
	// v.Previous = 0
	// v.Next = 0
	// v.Current = 0
	if version != "" {
		versionInt, err := strconv.Atoi(version)
		if err != nil {
			return v, err
		}
		v.Current = versionInt
		for i, id := range versionIds {
			if id == versionInt {
				if i > 0 {
					v.Previous = versionIds[i-1]
				}
				if i < len(versionIds)-1 {
					v.Next = versionIds[i+1]
				}
				break
			}
		}
	} else {
		if len(versionIds) > 0 {
			v.Previous = versionIds[len(versionIds)-1]
		}
	}

	return v, nil
}
