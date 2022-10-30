import cookie from 'cookie';
//import { COOKIENAME } from '../config.js';
import { COOKIENAME } from '$env/static/private';
import { getDb } from '$lib/server/db';

let db;

async function initDb() {
  db = await getDb();
}
initDb();

export async function handle({ event, resolve }) {

	event.locals.db = db;

  const cookies = cookie.parse(event.request.headers.cookie || '');
  //console.log(cookies)
  let r = null;
  if (cookies.hasOwnProperty(COOKIENAME) && db) {
    const c = await db.collection('sessions');
    r = await c.findOne({
      uuid: cookies[COOKIENAME]
    });
  }
  if (r) {
    event.locals.loggedIn = true;
    event.locals.user = r.user;
    event.locals.admin = r.admin;
    event.locals.sessionId = cookies[COOKIENAME];
  } else {
    event.locals.loggedIn = false;
  }

	return await resolve(event);
}

/* -> DEPR i think...? */
export function getSession(request) {

  if (request.locals.loggedIn) {
    return {
      loggedIn: true,
      user: request.locals.user,
      admin: request.locals.admin,
    }
  } else {
    return {
      loggedIn: false,
    }
  }

}
