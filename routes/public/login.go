package public

import (
	"bytes"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"

	"mypersonalwebsite/common"
	"mypersonalwebsite/config"
	"mypersonalwebsite/model"
)

// Path: "/login"
func RouteLogin(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	locals := common.GetLocals(r, db)
	var content bytes.Buffer
	err := templates["login"].Execute(&content, locals)
	if err != nil {
		common.ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
	common.RenderBaseTemplate(w, templates, "Login", &content, locals)
}

// Path: "/login"
// Method: POST
func RouteCheckLogin(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	// Get the username and password from the request
	err := r.ParseForm()
	if err != nil {
		common.ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		fmt.Println("Empty username or password")
		return
	}

	// Check if the username and password are valid
	var user model.User
	err = db.Get(&user, "SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}
		common.ErrorPage(w, err, http.StatusInternalServerError)
		return
	}

	// Check if the password is correct
	if !checkPasswordHash(password, user.PwHash) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate a random session ID
	sessionId, err := generateRandomString(32)
	if err != nil {
		// NOTE: at this point we're authenticated so let's see the error
		common.ErrorPage(w, err, http.StatusInternalServerError)
		return
	}

	// Store the session in the database
	_, err = db.Exec("INSERT INTO sessions (uuid, user_id, user_agent, ip, created_at) VALUES (?, ?, ?, ?, ?)",
		sessionId, user.Id, r.UserAgent(), r.RemoteAddr, time.Now().Format(time.RFC3339))
	if err != nil {
		common.SqlError(w, err)
		return
	}

	// Set the session cookie
	cookie := &http.Cookie{
		Name:  config.CookieName,
		Value: sessionId,
		// Expires:  time.Now().Add(30 * 24 * time.Hour),
		// MaxAge:   60 * 60 * 24 * 365 * 10, // 10 years
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	// Redirect to the dashboard
	http.Redirect(w, r, "/login", http.StatusFound)
}

func generateRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func checkPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
