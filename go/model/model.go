package model

import (
	"html/template"
)

type Entry struct {
	Id         int
	UserId     int `db:"user_id"`
	Type       string
	CreatedAt  string `db:"created_at"`
	ModifiedAt string `db:"modified_at"`
	Title      string
	Content    string
	Private    bool
}

type Link struct {
	Id         int
	UserId     int    `db:"user_id"`
	CreatedAt  string `db:"created_at"`
	ModifiedAt string `db:"modified_at"`
	Title      string
	Url        string
	Comment    string
	CategoryId string `db:"category_id"`
}

type LinkCategory struct {
	Id    int
	Name  string
	Links []Link
}

type LinkCategories struct {
	Categories []LinkCategory
}

type EntryPageData struct {
	Title      string
	Content    template.HTML
	ModifiedAt string
}

type LinkPageData struct {
}

type User struct {
	Id       int
	UserName string
	PwHash   string
}

type Session struct {
	Id        int
	Uuid      string
	UserId    int    `db:"user_id"`
	UserAgent string `db:"user_agent"`
	Ip        string
	CreatedAt string `db:"created_at"`
}

type Locals struct {
	LoggedIn bool
	UserName string
}
