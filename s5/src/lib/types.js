
/** @typedef {Object} ClientData
  * @property {boolean} loggedIn
  * @property {string | null} username
  */

/** @typedef {TimelineYearEntry|TimelineMonthEntry|TimelineDataEntry} TimelineEntry */

/** @typedef {Object} TimelineYearEntry
  * @property {'year'} type
  * @property {string} year
  * @property {null} month
  * @property {null} entry
  * @property {null} date
  */

/** @typedef {Object} TimelineMonthEntry
  * @property {'month'} type
  * @property {null} year
  * @property {string} month
  * @property {null} entry
  * @property {null} date
  */

/**
 * @typedef {Object} TimelineDataEntry
 * @property {'entry'} type
 * @property {null} year
 * @property {null} month
 * @property {import('$lib/server/db.js').EntryData} entry
 * @property {string} date
 */

export {};
