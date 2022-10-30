import cookie from 'cookie';
import { COOKIENAME } from '$env/static/private';

export async function post(request) {

  const error = {
    status: 401,
    body: {
      message: 'logout failed'
    }
  };

  if (!request.locals.loggedIn) {
    return error;
  }

  const db = request.locals.db;
  const c = await db.collection('sessions');
  const r = await c.deleteOne({
    uuid: request.locals.sessionId
  });
  if (!r) return error;

  return {
    body: '',
    headers: {
      'set-cookie': [
        cookie.serialize(
          COOKIENAME,
          request.locals.sessionId,
          {
            httpOnly: true,
            sameSite: true,
            //secure: true,
            maxAge: 0
          }
        )
      ]
    }
  };

}
