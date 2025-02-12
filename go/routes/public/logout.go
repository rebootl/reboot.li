package public

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"mypersonalwebsite/config"
)

func RouteLogout(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	// Get the session from the cookie
	cookie, err := r.Cookie(config.CookieName)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		fmt.Println("No cookie found")
		return
	}

	// Delete the session from the database
	_, err = db.Exec("DELETE FROM sessions WHERE uuid = ?", cookie.Value)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Delete the cookie
	cookie = &http.Cookie{
		Name:   config.CookieName,
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/login", http.StatusFound)
}
