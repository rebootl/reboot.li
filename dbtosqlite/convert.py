
import json
import sqlite3
# import os

class Entry:
    def __init__(self, type, date, topics, tags, deleted, version, last, private, text="", title="", comment="", **kwargs):
        self.type = type
        self.text = text
        self.date = date
        self.topics = topics
        self.tags = tags
        self.title = title
        self.comment = comment
        self.deleted = deleted
        self.version = version
        self.last = last
        self.private = private

def load_json_data(file_path):
    with open(file_path, 'r') as file:
        return json.load(file)

def insert_link_categories(db, topics):
    categories = set(topics)
    for category in categories:
        try:
            db.execute("INSERT OR IGNORE INTO link_categories (name) VALUES (?)", (category,))
        except sqlite3.Error as e:
            print(f"Error inserting category {category}: {e}")

def insert_links(db, link):
    if not link.topics:
        print(f"No topics provided for link: {link.title}")
        return

    try:
        category_id = db.execute("SELECT id FROM link_categories WHERE name = ?", (link.topics[0],)).fetchone()
        if category_id is None:
            print(f"Category not found for topics: {link.topics}")
            return

        result = db.execute(
            "INSERT INTO links (user_id, url, created_at, modified_at, title, comment, category_id) VALUES (?, ?, ?, ?, ?, ?, ?)",
            (1, link.text, link.date, link.date, link.title, link.comment, category_id[0])
        )
        db.commit()
        last_insert_id = result.lastrowid
        print(f"Inserted link with ID: {last_insert_id}")
    except sqlite3.Error as e:
        print(f"Error inserting link: {e}")

def insert_notes(db, entry):
    if not entry.last or entry.deleted:
        return

    note_type = "cheatsheet"
    if "cheat sheet" not in entry.tags:
        return

    try:
        result = db.execute(
            "INSERT INTO entries (type, user_id, content, created_at, modified_at, private) VALUES (?, ?, ?, ?, ?, ?)",
            (note_type, 1, entry.text, entry.date, entry.date, entry.private)
        )
        db.commit()
        last_insert_id = result.lastrowid
        print(f"Inserted note with ID: {last_insert_id}")
    except sqlite3.Error as e:
        print(f"Error inserting note: {e}")

def main():
    # Load your JSON data
    json_data = load_json_data("entries.json")
    entries = [Entry(**entry) for entry in json_data]

    # Initialize the SQLite database
    db = sqlite3.connect("db.sqlite")

    # Process entries
    for entry in entries:
        if entry.type == "link":
            insert_link_categories(db, entry.topics)
            insert_links(db, entry)
        elif entry.type == "note":
            insert_notes(db, entry)

    print("Data has been successfully inserted into the database.")
    db.close()

if __name__ == "__main__":
    main()
