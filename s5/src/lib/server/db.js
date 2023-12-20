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
  * @property {string} updated_at
  */

/**
  * @param {string} sessionId
  * @returns {SessionData | undefined}
  */
export function getSession(sessionId) {
  const stmt = db.prepare(`SELECT * FROM sessions
    JOIN users ON sessions.user_id = users.id
    WHERE sessions.id = ?`);
  const r = /** @type {SessionData | undefined} */ (stmt.get(sessionId));
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
