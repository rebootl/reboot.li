import { COOKIENAME } from '$env/static/private';
import { getSessionUser } from '$lib/server/db.js';

/**
  * @type {import('@sveltejs/kit').Handle}
  */
export const handle = ({ event, resolve }) => {

  const sessionId = event.cookies.get(COOKIENAME);

  /** @type {import('$lib/server/db').SessionUserData | undefined} */
  let r = undefined;
  if (sessionId) {
    r = getSessionUser(sessionId);
  }
  // console.log('r', r);

  if (r) {
    event.locals = {
      user: {
        name: r.username,
      }
    };
  } else {
    event.locals = {
      user: null,
    };
  }

  const response = resolve(event);
  return response;
};
