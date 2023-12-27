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
  */

/**
  * @param {string} sessionId
  * @returns {SessionUserData | undefined}
  */
export function getSessionUser(sessionId) {
  const stmt = db.prepare(`SELECT users.username AS username FROM sessions
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
  */
export function createSession(data) {
  console.log('createSession', data);
  const stmt = db.prepare(`INSERT INTO sessions (uuid, user_id, user_agent, ip, created_at)
    VALUES (?, ?, ?, ?, datetime('now'))`);
  const r = stmt.run(data.uuid, data.userId, data.userAgent, data.host);
  console.log('createSession r', r);
  return r;
}
