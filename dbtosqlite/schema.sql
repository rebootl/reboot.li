BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "users" (
	"id"	INTEGER NOT NULL UNIQUE,
	"username"	TEXT NOT NULL UNIQUE,
	"pwhash"	TEXT NOT NULL,
	"created_at"	TEXT,
	"modified_at"	TEXT,
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "sessions" (
	"id"	INTEGER NOT NULL UNIQUE,
	"user_id"	INTEGER NOT NULL,
	"uuid"	TEXT NOT NULL UNIQUE,
	"user_agent"	TEXT,
	"ip"	TEXT,
	"created_at"	TEXT NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "entry_to_tag" (
	"id"	INTEGER NOT NULL UNIQUE,
	"entry_id"	INTEGER NOT NULL,
	"tag_id"	INTEGER NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT),
	FOREIGN KEY("tag_id") REFERENCES "entry_tags"("id"),
	FOREIGN KEY("entry_id") REFERENCES "entries"("id")
);
CREATE TABLE IF NOT EXISTS "link_to_tag" (
	"id"	INTEGER NOT NULL UNIQUE,
	"link_id"	INTEGER NOT NULL,
	"tag_id"	INTEGER NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT),
	FOREIGN KEY("tag_id") REFERENCES "link_tags"("id"),
	FOREIGN KEY("link_id") REFERENCES "link"("id")
);
CREATE TABLE IF NOT EXISTS "entry_tags" (
	"id"	INTEGER NOT NULL UNIQUE,
	"user_id"	INTEGER NOT NULL,
	"name"	TEXT NOT NULL UNIQUE,
	"color"	TEXT,
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "link_tags" (
	"id"	INTEGER NOT NULL UNIQUE,
	"user_id"	INTEGER NOT NULL,
	"name"	TEXT NOT NULL UNIQUE,
	"color"	TEXT,
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "images" (
	"id"	INTEGER NOT NULL UNIQUE,
	"entry_id"	INTEGER NOT NULL,
	"user_id"	INTEGER NOT NULL,
	"path"	TEXT NOT NULL UNIQUE,
	"comment"	TEXT,
	"created_at"	TEXT NOT NULL,
	"preview_data"	TEXT,
	FOREIGN KEY("user_id") REFERENCES "users"("id"),
	PRIMARY KEY("id" AUTOINCREMENT),
	FOREIGN KEY("entry_id") REFERENCES "entries"("id")
);
CREATE TABLE IF NOT EXISTS "link_categories" (
	"id"	INTEGER NOT NULL UNIQUE,
	"name"	TEXT UNIQUE,
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "entries" (
	"id"	INTEGER NOT NULL UNIQUE,
	"user_id"	INTEGER NOT NULL,
	"type"	TEXT NOT NULL,
	"created_at"	TEXT NOT NULL,
	"modified_at"	TEXT,
	"title"	TEXT,
	"content"	TEXT,
	"private"	INTEGER NOT NULL,
	FOREIGN KEY("user_id") REFERENCES "users"("id"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "links" (
	"id"	INTEGER NOT NULL UNIQUE,
	"user_id"	INTEGER NOT NULL,
	"created_at"	TEXT NOT NULL,
	"modified_at"	TEXT,
	"title"	TEXT,
	"url"	TEXT,
	"comment"	TEXT,
	"category_id"	INTEGER NOT NULL,
	FOREIGN KEY("user_id") REFERENCES "users"("id"),
	FOREIGN KEY("category_id") REFERENCES "link_categories"("id"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
COMMIT;
PRAGMA journal_mode = WAL;
VACUUM;
