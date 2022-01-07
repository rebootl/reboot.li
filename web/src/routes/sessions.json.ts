import cookie from 'cookie';
import { COOKIENAME } from '../../config.js';
import { ObjectId } from 'mongodb';

export async function get(request) {

  if (!request.locals.loggedIn)
    return { status: 403 };

  const db = request.locals.db;

  const c = await db.collection('sessions');

  const q = { user: request.locals.user };

  const r = await c.find(q).sort({ createdAt: -1 }).toArray();
  if (!r) return { status: 404 };

  const cookies = cookie.parse(request.headers.cookie || '');
  const uuid = cookies[COOKIENAME];

  for (const s of r) {
    if (s.uuid === uuid) {
      s.current = true;
    } else {
      s.current = false;
    }
    delete s.uuid;
  }

  return {
    body: r
  };
}

export async function del(request) {
  //console.log(request)

  const b = request.body;

  // user logged in and username in entry
  if (!request.locals.loggedIn || request.locals?.user !== b.user)
    return { status: 403 }

  const db = request.locals.db;
  const c = await db.collection('sessions');
  const r = await c.deleteOne({ _id: new ObjectId(b._id), user: b.user });
  if (!r?.deletedCount) return { status: 400 };

  return {
    body: {
      success: true,
      result: r
    }
  };
}
