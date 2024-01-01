CREATE TABLE sqlite_sequence(name,seq);
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
	"uuid"	TEXT NOT NULL UNIQUE,
	"user_id"	INTEGER NOT NULL,
	"user_agent"	TEXT,
	"ip"	TEXT,
	"created_at"	TEXT NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "entry_to_tag" (
	"id"	INTEGER NOT NULL UNIQUE,
	"entry_id"	INTEGER NOT NULL,
	"tag_id"	INTEGER NOT NULL,
	FOREIGN KEY("entry_id") REFERENCES "entries"("id"),
	FOREIGN KEY("tag_id") REFERENCES "tags"("id"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "tags" (
	"id"	INTEGER NOT NULL UNIQUE,
	"user_id"	INTEGER NOT NULL,
	"name"	TEXT NOT NULL UNIQUE,
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "entries" (
	"id"	INTEGER NOT NULL UNIQUE,
	"type"	INTEGER NOT NULL,
	"user_id"	INTEGER NOT NULL,
	"content"	TEXT,
	"created_at"	TEXT NOT NULL,
	"modified_at"	TEXT,
	"version"	INTEGER NOT NULL,
	"current"	INTEGER NOT NULL,
	"title"	TEXT,
	"comment"	TEXT,
	"pinned"	INTEGER,
	"private"	INTEGER,
	FOREIGN KEY("user_id") REFERENCES "users"("id"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
