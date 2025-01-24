import { Database } from 'bun:sqlite';
import fs from 'fs';

// Load your JSON data
const jsonData = JSON.parse(fs.readFileSync('entries.json', 'utf-8'));

// Initialize the SQLite database
const db = new Database('db.sqlite');

// Function to insert link categories
function insertLinkCategories(topics) {
    const categories = new Set(topics);
    for (const category of categories) {
        db.query('INSERT OR IGNORE INTO link_categories (name) VALUES (?)', [category]);
    }
}

// Function to insert links
function insertLinks(link) {
    const { title, comment, text, user, date, topics } = link;

    // Use .get() to retrieve the category ID
    const categoryResult = db.query('SELECT id FROM link_categories WHERE name = ?', [topics[0]]).get();
    const categoryId = categoryResult ? categoryResult.id : null;

    if (categoryId) {
        const result = db.query(
            'INSERT INTO links (user_id, url, created_at, modified_at, title, comment, category_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)',
            [1, text, date, date, title, comment, categoryId]
        ).run();
        console.log(`Inserted link with ID: ${result.lastInsertRowid}`);
    } else {
        console.warn(`Category not found for topics: ${topics}`);
    }
}

// Function to insert notes
function insertNotes(note) {
    const { text, user, date, tags } = note;
    const type = tags.includes('cheat sheet') ? 'cheatsheet' : 'note';

    if (type) {
        console.log(type)
        const result = db.query(
            'INSERT INTO entries (type, user_id, content, created_at, modified_at) VALUES (:type, :user, :text, :date, :date)',
        ).run({type, user: 1, text, date, date});
        console.log(`Inserted note with ID: ${result.lastInsertRowid}`);
    }
}
// Main function to process the JSON data
function processEntries() {
    // Check if entries exist in the JSON data
    const entries = jsonData || []; // Adjust this line based on your actual JSON structure

    for (const entry of entries) {
        if (entry.type === 'link') {
            insertLinkCategories(entry.topics);
            insertLinks(entry);
        } else if (entry.type === 'note') {
            insertNotes(entry);
        }
        // You can add more conditions here for other types if needed
    }
}

// Run the processing function
processEntries()
    .then(() => {
        console.log('Data has been successfully inserted into the database.');
    })
    .catch((error) => {
        console.error('Error inserting data:', error);
    })
    .finally(() => {
        db.close();
    });
