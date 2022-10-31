import { json, error } from '@sveltejs/kit';
import { COOKIENAME } from '$env/static/private';

export async function POST({ request, locals, cookies }) {

  const error = {
    status: 401,
    body: {
      message: 'logout failed'
    }
  };

  if (!locals.loggedIn) {
    throw error(401, 'logout failed');
  }

  const db = locals.db;
  const c = await db.collection('sessions');
  const r = await c.deleteOne({
    uuid: locals.sessionId
  });
  if (!r) throw error(401, 'logout failed');

  cookies.delete(COOKIENAME);

  return json('');
}
