package model

import (
	"database/sql"
	"fmt"
	"html/template"
	"time"

	"github.com/jmoiron/sqlx"
)

// database models

type Entry struct {
	Id         int
	UserId     int `db:"user_id"`
	Type       string
	CreatedAt  string `db:"created_at"`
	ModifiedAt string `db:"modified_at"`
	Title      string
	Content    string
	Private    bool
	Tags       []Tag // not actually in the entry table FIXME where is this used?
}

type EntryVersion struct {
	Id        int
	EntryId   int `db:"entry_id"`
	Title     string
	Content   string
	CreatedAt string `db:"created_at"`
}

type Tag struct {
	Id     int
	UserId int `db:"user_id"`
	Name   string
	Color  string
}

type Link struct {
	Id         int
	UserId     int    `db:"user_id"`
	CreatedAt  string `db:"created_at"`
	ModifiedAt string `db:"modified_at"`
	Title      string
	Url        string
	Comment    string
	CategoryId int `db:"category_id"`
}

type LinkCategory struct {
	Id    int
	Name  string
	Links []Link // not actually in the entry table FIXME where is this used?
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

// page data models, these are used to pass data to the templates

type LinkCategories struct {
	Categories []LinkCategory
}

type EntryPageData struct {
	Id         int
	Title      string
	Content    template.HTML
	ModifiedAt string
	Tags       []Tag
	VersionIds []int
	IsVersion  bool
	Locals
}

type EditPageData struct {
	Entry      Entry
	Title      string
	Type       string
	ModifiedAt string
	AllTags    []TagWithStatus
	Ref        string
}

type EditLinkPageData struct {
	Link          Link
	Title         string
	ModifiedAt    string
	AllCategories []LinkCategory
	Ref           string
}

type TagWithStatus struct {
	Tag      Tag
	Selected bool
}

type ListPageData struct {
	Entries []Entry
	Motd    string
	Locals
}

type BasePageData struct {
	Title   string
	Content template.HTML
	Locals
}

type Locals struct {
	LoggedIn bool
	UserName string
}

// database functions

func GetEntryById(db *sqlx.DB, locals Locals, id string) (Entry, error) {
	var q string
	if locals.LoggedIn {
		q = "SELECT * FROM entries WHERE id = ?"
	} else {
		q = "SELECT * FROM entries WHERE id = ? AND private = 0"
	}

	var entry Entry
	err := db.Get(&entry, q, id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
			// TODO: return a 404 page
		} else {
			fmt.Println(err)
		}
		return entry, err
	}

	tags, err := GetTagsByEntryId(db, id)
	if err != nil {
		fmt.Println(err)
		return entry, err
	}
	entry.Tags = tags
	return entry, nil
}

func GetTagsByEntryId(db *sqlx.DB, id string) ([]Tag, error) {
	var tags []Tag
	err := db.Select(&tags, `SELECT t.id, t.name, t.color
		FROM entry_tags t
		JOIN entry_to_tag et ON t.id = et.tag_id
		WHERE et.entry_id = ?`, id)
	return tags, err
}

func GetAllEntryTags(db *sqlx.DB) ([]Tag, error) {
	var tags []Tag
	err := db.Select(&tags, `SELECT * FROM entry_tags ORDER BY name`)
	return tags, err
}

func UpdateEntryTags(db *sqlx.DB, id string, selectedTagNames []string) error {

	existingTags, err := GetTagsByEntryId(db, id)
	if err != nil {
		return err
	}

	selectedTagsByNames := make(map[string]bool)
	for _, tagName := range selectedTagNames {
		selectedTagsByNames[tagName] = true
	}

	for _, tag := range existingTags {
		if !selectedTagsByNames[tag.Name] {
			_, err = db.Exec(
				"DELETE FROM entry_to_tag WHERE entry_id = ? AND tag_id = ?",
				id, tag.Id,
			)
			if err != nil {
				return err
			}
		}
	}

	allTags, err := GetAllEntryTags(db)
	if err != nil {
		return err
	}
	allTagsByNames := make(map[string]Tag)
	for _, tag := range allTags {
		allTagsByNames[tag.Name] = tag
	}

	existingTagNames := make(map[string]bool)
	for _, tag := range existingTags {
		existingTagNames[tag.Name] = true
	}

	for _, tagName := range selectedTagNames {
		if !existingTagNames[tagName] {
			_, err = db.Exec(
				"INSERT INTO entry_to_tag (entry_id, tag_id) VALUES (?, ?)",
				id, allTagsByNames[tagName].Id,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func GetEntryTagById(db *sqlx.DB, id string) (Tag, error) {
	var tag Tag
	err := db.Get(&tag, "SELECT * FROM entry_tags WHERE id = ?", id)
	return tag, err
}

func GetLinkById(db *sqlx.DB, id string) (Link, error) {
	var link Link
	err := db.Get(&link, "SELECT * FROM links WHERE id = ?", id)
	return link, err
}

func GetAllLinkCategories(db *sqlx.DB) ([]LinkCategory, error) {
	var categories []LinkCategory
	err := db.Select(&categories, `SELECT * FROM link_categories ORDER BY name`)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func GetLinkCategoryById(db *sqlx.DB, id string) (LinkCategory, error) {
	var category LinkCategory
	err := db.Get(&category, "SELECT * FROM link_categories WHERE id = ?", id)
	if err != nil {
		return category, err
	}
	return category, err
}

func SaveVersion(db *sqlx.DB, locals Locals, id string) error {
	// get the entry
	entry, err := GetEntryById(db, locals, id)
	if err != nil {
		return err
	}

	timestamp := time.Now().Format(time.RFC3339)
	_, err = db.Exec(`
			INSERT INTO entries_versions (entry_id, title, content, created_at)
			VALUES ($1, $2, $3, $4)
			`, id, entry.Title, entry.Content, timestamp)
	if err != nil {
		return err
	}

	return nil
}

func GetVersionIds(db *sqlx.DB, id int) ([]int, error) {
	var versionIds []int
	err := db.Select(&versionIds, "SELECT id FROM entries_versions WHERE entry_id = ? ORDER BY id ASC", id)
	if err != nil {
		return nil, err
	}
	return versionIds, nil
}

func GetEntryVersion(db *sqlx.DB, id int, versionId string) (EntryVersion, error) {
	var version EntryVersion
	err := db.Get(&version, "SELECT * FROM entries_versions WHERE entry_id = ? AND id = ?", id, versionId)
	if err != nil {
		return version, err
	}
	return version, nil
}
