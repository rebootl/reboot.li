package private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"regexp"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"mypersonalwebsite/common"
	"mypersonalwebsite/model"
)

// Path: "/edit-link/{id}"
func RouteEditLink(
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

	var (
		link  model.Link
		title string
		err   error
	)

	if vars["id"] == "new" {
		link = model.Link{
			Id:         0,
			Title:      "",
			Comment:    "",
			ModifiedAt: "",
		}
		title = "New Entry"
	} else {
		link, err = model.GetLinkById(db, vars["id"])
		if err != nil {
			common.SqlError(w, err)
			return
		}
		title = "Edit Entry"
	}

	allCategories, err := model.GetAllLinkCategories(db)
	if err != nil {
		common.SqlError(w, err)
		return
	}

	// preprocesse date
	modifiedAt, _ := time.Parse(time.RFC3339, link.ModifiedAt)

	var content bytes.Buffer
	err = templates["edit-link"].Execute(&content, model.EditLinkPageData{
		Link:          link,
		ModifiedAt:    modifiedAt.Format("2006-01-02 15:04h"),
		Title:         title,
		AllCategories: allCategories,
	})
	if err != nil {
		common.ErrorPage(w, err, http.StatusInternalServerError)
		return
	}

	common.RenderBaseTemplate(w, templates, title, &content, locals)
}

// Path: "/update-link"
// Method: POST
func RouteUpdateLink(
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
	url := r.FormValue("url")
	title := r.FormValue("title")
	comment := r.FormValue("comment")
	category_id := r.FormValue("categories")
	fmt.Println(url)

	// Validate the form data
	if title == "" || url == "" {
		http.Error(w, "Title and url are required", http.StatusBadRequest)
		return
	}

	timestamp := time.Now().Format(time.RFC3339)
	if id == "0" {
		// Insert a new entry into the database
		_, err = db.Exec(`
			INSERT INTO links (url, title, comment, category_id, created_at, modified_at, user_id)
				VALUES ($1, $2, $3, $4, $5, $5, $6)
		`, url, title, comment, category_id, timestamp, 1)
	} else {
		// Update the entry in the database
		_, err = db.Exec(`
			UPDATE links
			SET url = $1, title = $2, comment = $3, category_id = $4, modified_at = $5
			WHERE id = $6
		`, url, title, comment, category_id, timestamp, id)
	}
	if err != nil {
		common.SqlError(w, err)
		return
	}

	http.Redirect(w, r, "/links", 302)
}

// Path: "/delete-link"
// Method: POST
func RouteDeleteLink(
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

	// delete link to tag links first
	_, err = db.Exec(`
		DELETE FROM link_to_tag
		WHERE link_id = $1
	`, id)
	if err != nil {
		common.SqlError(w, err)
		return
	}

	_, err = db.Exec(`
		DELETE FROM links
		WHERE id = $1
	`, id)
	if err != nil {
		common.SqlError(w, err)
		return
	}

	http.Redirect(w, r, "/links", 302)
}

// Path: "/api/get-title/{url}"
// Method: GET
func RouteGetTitle(
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

	url := r.URL.Query().Get("url")

	fmt.Println(url)
	if url == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		fmt.Println("URL is required")
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error requesting page: "+err.Error(), 255)
		fmt.Println("Error requesting page")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Error requesting page: "+resp.Status, 255)
		fmt.Println("Error requesting page: " + resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading page: "+err.Error(), 255)
		fmt.Println("Error reading page")
		return
	}

	title, err := getTitle(string(body))
	if err != nil {
		http.Error(w, "Error getting title: "+err.Error(), 255)
		fmt.Println("Error getting title")
		return
	}

	jsonResponse, err := json.Marshal(struct {
		Title string `json:"title"`
	}{Title: title})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func getTitle(html string) (string, error) {
	re := regexp.MustCompile(`<title>(.*?)</title>`)
	matches := re.FindAllStringSubmatch(html, -1)

	if len(matches) == 0 {
		return "", fmt.Errorf("title not found")
	}

	return matches[0][1], nil
}
