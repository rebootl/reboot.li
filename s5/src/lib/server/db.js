import Database from 'better-sqlite3';
import { DBPATH } from '$env/static/private'; 

export const db = new Database(DBPATH);

/**
  * @typedef {Object} SessionData
  * @property {number} id
  * @property {string} uuid
  * @property {number} user_id
  * @property {string} user_agent
  * @property {string} ip
  * @property {string} created_at
  */

/**
  * @typedef {Object} SessionUserData
  * @property {string} username
  * @property {number} user_id
  */

/**
  * @param {string} sessionId
  * @returns {SessionUserData | undefined}
  */
export function getSessionUser(sessionId) {
  const stmt = db.prepare(`SELECT users.username AS username, sessions.user_id AS user_id FROM sessions
    JOIN users ON sessions.user_id = users.id
    WHERE sessions.uuid = ?`);
  const r = /** @type {SessionUserData | undefined} */ (stmt.get(sessionId));
  return r;
}

/**
  * @typedef {Object} UserData
  * @property {number} id
  * @property {string} username
  * @property {string} pwhash
  * @property {string} created_at
  * @property {string} updated_at
  */

/**
  * @param {string} username
  * @returns {UserData | undefined}
  */
export function getUser(username) {
  const stmt = db.prepare(`SELECT * FROM users WHERE username = ?`);
  const r = /** @type {UserData | undefined} */ (stmt.get(username));
  return r;
}

/**
  * @typedef {Object} SessionCreateData
  * @property {string} uuid
  * @property {number} userId
  * @property {string} userAgent
  * @property {string} host
  */

/**
  * @param {SessionCreateData} data
  * @returns {Database.RunResult}
  */
export function createSession(data) {
  // console.log('createSession', data);
  const stmt = db.prepare(`INSERT INTO sessions (uuid, user_id, user_agent, ip, created_at)
    VALUES (?, ?, ?, ?, datetime('now'))`);
  const r = stmt.run(data.uuid, data.userId, data.userAgent, data.host);
  // console.log('createSession r', r);
  return r;
}

/**
  * @param {string} sessionId
  * @returns {Database.RunResult}
  */
export function destroySession(sessionId) {
  const stmt = db.prepare(`DELETE FROM sessions WHERE uuid = ?`);
  const r = stmt.run(sessionId);
  return r;
}

/**
  * @typedef {Object} CreateEntryData
  * @property {number} userId
  * @property {string} type
  * @property {string} title
  * @property {string} content
  * @property {string} comment
  * @property {string} private
  * @property {string} pinned
  * @property {string[]} tags
  */

/**
  * @param {CreateEntryData} data
  * @returns {Database.RunResult}
  */
export function createEntry(data) {
  const stmt = db.prepare(`INSERT INTO entries (user_id, type, title, content, comment, private, pinned, created_at)
    VALUES (?, ?, ?, ?, ?, ?, ?, datetime('now'))`);
  const r = stmt.run(data.userId, data.type, data.title, data.content, data.comment, data.private, data.pinned);

  for (const t of data.tags) {
    const stmt2 = db.prepare(`INSERT INTO tags (user_id, entry_id, name, created_at)
      VALUES (?, ?, ?, datetime('now'))`);
    const r2 = stmt2.run(data.userId, r.lastInsertRowid, t);
  }
  return r;
}

/**
  * @typedef {Object} Tag
  * @property {number} id
  * @property {number} user_id
  * @property {string} name
*/

/**
  * @typedef {Object} EntryData
  * @property {number} id
  * @property {number} user_id
  * @property {string} type
  * @property {string} title
  * @property {string} content
  * @property {string} comment
  * @property {string} private
  * @property {string} pinned
  * @property {string} created_at
  * @property {string} updated_at
  * @property {number} version
  * @property {number} last
  * @property {Tag[]} tags
  */

/**
  * @param {number} userId
  * @param {number} entryId
  * @param {boolean} loggedIn
  * @returns {EntryData | undefined}
  */
export function getEntry(userId, entryId, loggedIn = false) {
  let stmt;
  if (!loggedIn) {
    stmt = db.prepare(`SELECT * FROM entries WHERE user_id = ? AND id = ? AND private = 0 AND current = 1`);
  } else {
    stmt = db.prepare(`SELECT * FROM entries WHERE user_id = ? AND id = ? AND current = 1`);
  }
  const r = /** @type {EntryData | undefined} */ (stmt.get(userId, entryId));
  // get tags
  if (r) {
    const stmt2 = db.prepare(`SELECT * FROM tags WHERE user_id = ? AND entry_id = ?`);
    const tags = /** @type {Tag[]} */ (stmt2.all(userId, entryId));
    r.tags = tags;
  }
  return r;
}

/**
  * @param {number} userId
  * @param {string|null} type
  * @param {boolean} loggedIn
  * @param {number} limit
  * @param {number} offset
  * @returns {EntryData[]}
  */
export function getEntries(userId, type = null, loggedIn = false, limit = 99, offset = 0) {
  let stmt;
  if (type === null) {
    if (!loggedIn) {
      stmt = db.prepare(`SELECT * FROM entries WHERE user_id = ? AND private = 0 ORDER BY created_at DESC LIMIT ? OFFSET ?`);
    } else {
      stmt = db.prepare(`SELECT * FROM entries WHERE user_id = ? ORDER BY created_at DESC LIMIT ? OFFSET ?`);
    }
  } else {
    if (!loggedIn) {
      stmt = db.prepare(`SELECT * FROM entries WHERE user_id = ? AND type = ? AND private = 0 ORDER BY created_at DESC LIMIT ? OFFSET ?`);
    } else {
      stmt = db.prepare(`SELECT * FROM entries WHERE user_id = ? AND type = ? ORDER BY created_at DESC LIMIT ? OFFSET ?`);
    }
  }
  const r = /** @type {EntryData[]} */ (stmt.all(userId, type, limit, offset));
  return r;
}
