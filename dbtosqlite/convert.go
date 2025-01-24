package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Entry struct {
	Type    string   `json:"type"`
	Text    string   `json:"text"`
	Date    string   `json:"date"`
	Topics  []string `json:"topics"`
	Tags    []string `json:"tags"`
	Title   string   `json:"title"`
	Comment string   `json:"comment"`
	Deleted bool     `json:"deleted"`
	Version int      `json:"version"`
	Last    bool     `json:"last"`
	Private bool     `json:"private"`
}

func main() {
	// Load your JSON data
	jsonData, err := os.ReadFile("entries.json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	var entries []Entry
	if err := json.Unmarshal(jsonData, &entries); err != nil {
		log.Fatalf("Error parsing JSON data: %v", err)
	}

	// Initialize the SQLite database
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Process entries
	for _, entry := range entries {
		if entry.Private == true {
			continue
		}
		if entry.Type == "link" {
			insertLinkCategories(db, entry.Topics)
			insertLinkTags(db, entry.Tags)
			insertLinks(db, entry)
		} else if entry.Type == "note" {
			insertEntryTags(db, entry.Tags)
			insertNotes(db, entry)
		}
	}

	fmt.Println("Data has been successfully inserted into the database.")
}

func insertLinkCategories(db *sql.DB, topics []string) {
	categories := make(map[string]struct{})
	for _, category := range topics {
		categories[category] = struct{}{}
	}

	for category := range categories {
		_, err := db.Exec("INSERT OR IGNORE INTO link_categories (name) VALUES (?)", category)
		if err != nil {
			log.Printf("Error inserting category %s: %v", category, err)
		}
	}
}

func insertLinkTags(db *sql.DB, tags []string) {
	// fmt.Println("Inserting tags: ", tags)
	for _, tag := range tags {
		_, err := db.Exec("INSERT OR IGNORE INTO link_tags (user_id, name) VALUES (?, ?)", 1, tag)
		if err != nil {
			log.Printf("Error inserting tag %s: %v", tag, err)
		}
	}
}

func insertEntryTags(db *sql.DB, tags []string) {
	for _, tag := range tags {
		_, err := db.Exec("INSERT OR IGNORE INTO entry_tags (user_id, name) VALUES (?, ?)", 1, tag)
		if err != nil {
			log.Printf("Error inserting tag %s: %v", tag, err)
		}
	}
}

func insertLinks(db *sql.DB, link Entry) {
	if len(link.Topics) == 0 {
		log.Printf("No topics provided for link: %s", link.Title)
		return
	}

	var categoryId int
	err := db.QueryRow("SELECT id FROM link_categories WHERE name = ?", link.Topics[0]).Scan(&categoryId)
	if err != nil {
		log.Printf("Category not found for topics: %v", link.Topics)
		return
	}

	result, err := db.Exec(
		"INSERT INTO links (user_id, url, created_at, modified_at, title, comment, category_id) VALUES (?, ?, ?, ?, ?, ?, ?)",
		1, link.Text, link.Date, link.Date, link.Title, link.Comment, categoryId,
	)
	if err != nil {
		log.Printf("Error inserting link: %v", err)
		return
	}

	lastInsertID, _ := result.LastInsertId()
	// for every tag find the tag id and inserted into link_to_tag table
	for _, tag := range link.Tags {
		var tagId int
		err := db.QueryRow("SELECT id FROM link_tags WHERE name = ?", tag).Scan(&tagId)
		if err != nil {
			log.Printf("Tag not found: %v", tag)
			continue
		}
		_, err = db.Exec(
			"INSERT INTO link_to_tag (link_id, tag_id) VALUES (?, ?)",
			lastInsertID, tagId,
		)
		if err != nil {
			log.Printf("Error inserting link to tag: %v", err)
		}
	}

	fmt.Printf("Inserted link with ID: %d\n", lastInsertID)
}

func insertNotes(db *sql.DB, entry Entry) {
	if !entry.Last || entry.Deleted {
		return
	}
	noteType := "cheatsheet"
	if !contains(entry.Tags, "cheat sheet") {
		return
	}

	// extract the 1st line of the content and strip any # characters
	lines := strings.Split(entry.Text, "\n")
	firstLine := strings.TrimPrefix(lines[0], "#")
	firstLine = strings.TrimSpace(firstLine)

	updatedContent := strings.Join(lines[1:], "\n")
	updatedContent = strings.Trim(updatedContent, "\n")

	result, err := db.Exec(
		"INSERT INTO entries (type, user_id, title, content, created_at, modified_at, private) VALUES (?, ?, ?, ?, ?, ?, ?)",
		noteType, 1, firstLine, updatedContent, entry.Date, entry.Date, entry.Private,
	)
	if err != nil {
		log.Printf("Error inserting note: %v", err)
		return
	}

	lastInsertID, _ := result.LastInsertId()
	// for every tag find the tag id and inserted into entry_to_tag table
	for _, tag := range entry.Tags {
		var tagId int
		err := db.QueryRow("SELECT id FROM entry_tags WHERE name = ?", tag).Scan(&tagId)
		if err != nil {
			log.Printf("Tag not found: %v", tag)
			continue
		}
		_, err = db.Exec(
			"INSERT INTO entry_to_tag (entry_id, tag_id) VALUES (?, ?)",
			lastInsertID, tagId,
		)
		if err != nil {
			log.Printf("Error inserting entry to tag: %v", err)
		}
	}

	fmt.Printf("Inserted note with ID: %d\n", lastInsertID)
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
