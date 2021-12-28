import cookie from 'cookie';
import { COOKIENAME } from '../config.js';
import { getDb } from '$lib/db';

let db;

async function initDb() {
  db = await getDb();
}
initDb();

export async function handle({ request, resolve }) {

	request.locals.db = db;

  const cookies = cookie.parse(request.headers.cookie || '');
  //console.log(cookies)
  let r = null;
  if (cookies.hasOwnProperty(COOKIENAME) && db) {
    const c = await db.collection('sessions');
    r = await c.findOne({
      uuid: cookies[COOKIENAME]
    });
  }
  if (r) {
    request.locals.loggedIn = true;
    request.locals.user = r.user;
    request.locals.admin = r.admin;
    request.locals.sessionId = cookies[COOKIENAME];
  } else {
    request.locals.loggedIn = false;
  }

	return await resolve(request);
}

export function getSession(request) {

  if (request.locals.loggedIn) {
    return {
      loggedIn: true,
      user: request.locals.user,
      admin: request.locals.admin,
      sessionId: request.locals.sessionId,
    }
  } else {
    return {
      loggedIn: false,
    }
  }

}
