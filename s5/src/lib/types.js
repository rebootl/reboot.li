
/** @typedef {Object} ClientData
  * @property {boolean} loggedIn
  * @property {string | null} username
  */

/** @typedef {Object} TimelineEntry
  * @property {'year' | 'month' | 'entry'} type
  * @property {string | null} year
  * @property {string | null} month
  * @property {import('$lib/server/db.js').EntryData | null} entry
  * @property {string | null} date
  */

export {};
