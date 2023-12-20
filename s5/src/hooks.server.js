import { COOKIENAME } from '$env/static/private';
import { getSession } from '$lib/server/db.js';

// import type { Handle } from '@sveltejs/kit';
/**
  * @type {import('@sveltejs/kit').Handle}
  */
export const handle = (async ({ event, resolve }) => {

  const sessionId = event.cookies.get(COOKIENAME);

  /** @type {import('$lib/server/db').SessionData | undefined} */
  let r = undefined;
  if (sessionId) {
    r = getSession(sessionId);
  }
  console.log('r', r);

  if (r) {
    event.locals = {
      user: {
        // name: r.username,
        id: r.user_id,
      }
    };
  } else {
    event.locals = {
      user: null,
    };
  }

  const response = await resolve(event);
  return response;
});
