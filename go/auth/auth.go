package auth

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"mypersonalwebsite/config"
	"mypersonalwebsite/model"
)

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
