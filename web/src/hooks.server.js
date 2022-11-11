import { COOKIENAME } from '$env/static/private';
import { getDb } from '$lib/server/db';

let db;

async function initDb() {
  db = await getDb();
}
initDb();

export async function handle({ event, resolve }) {

  //console.log('hook handle')
	event.locals.db = db;

  const sessionId = event.cookies.get(COOKIENAME) || null;
  //console.log(sessionId)
  let r = null;
  if (sessionId && db) {
    const c = await db.collection('sessions');
    r = await c.findOne({
      uuid: sessionId
    });
  }
  if (r) {
    event.locals.loggedIn = true;
    event.locals.user = r.user;
    event.locals.admin = r.admin;
    event.locals.sessionId = sessionId;
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
