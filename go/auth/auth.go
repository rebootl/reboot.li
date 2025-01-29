package auth

import (
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

	"mypersonalwebsite/config"
	"mypersonalwebsite/model"
)

func GetLocals(r *http.Request, db *sqlx.DB) model.Locals {
	// Check if the user is logged in
	cookie, err := r.Cookie(config.CookieName)
	if err != nil {
		// fmt.Println("No cookie found")
		return model.Locals{LoggedIn: false, UserName: ""}
	}

	// Get the session from the database
	var session model.Session
	err = db.Get(&session, "SELECT * FROM sessions WHERE uuid = ?", cookie.Value)
	if err != nil {
		// fmt.Println("No session found")
		return model.Locals{LoggedIn: false, UserName: ""}
	}

	// Get the user from the database
	var user model.User
	err = db.Get(&user, "SELECT * FROM users WHERE id = ?", session.UserId)
	if err != nil {
		// fmt.Println("No user found")
		return model.Locals{LoggedIn: false, UserName: ""}
	}

	return model.Locals{LoggedIn: true, UserName: user.UserName}
}

func RouteCheckLogin(
	entryType string,
	w http.ResponseWriter,
	r *http.Request,
	db *sqlx.DB,
	templates map[string]*template.Template,
) {
	// Get the username and password from the request
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Check if the username and password are valid
	var user model.User
	err := db.Get(&user, "SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if the password is correct
	if !checkPasswordHash(password, user.PwHash) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate a random session ID
	sessionID, err := generateRandomString(32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Store the session in the database
	_, err = db.Exec("INSERT INTO sessions (id, uuid, user_id, user_agent, ip, created_at) VALUES (NULL, ?, ?, ?, ?, ?)",
		sessionID, user.Id, r.UserAgent(), r.RemoteAddr, time.Now().Format(time.RFC3339))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the session cookie
	cookie := &http.Cookie{
		Name:  config.CookieName,
		Value: sessionID,
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
		fmt.Println("No cookie found")
		return
	}

	// Delete the session from the database
	_, err = db.Exec("DELETE FROM sessions WHERE uuid = ?", cookie.Value)
	if err != nil {
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
