//import cookie from 'cookie';
import { json, error } from '@sveltejs/kit';
import { COOKIENAME } from '$env/static/private';
import { ObjectId } from 'mongodb';

export async function GET({ request, params, locals, cookies }) {

  if (!locals.loggedIn)
    throw error(403, 'Not allowed');

  const db = locals.db;

  const c = await db.collection('sessions');

  const q = { user: locals.user };

  const r = await c.find(q).sort({ createdAt: -1 }).toArray();
  if (!r) throw error(404, 'Not found');

  //const cookies = cookie.parse(request.headers.cookie || '');
  const uuid = cookies.get(COOKIENAME) || null;

  for (const s of r) {
    if (s.uuid === uuid) {
      s.current = true;
    } else {
      s.current = false;
    }
    delete s.uuid;
  }

  return json(r);
}

export async function DELETE({ request, locals }) {
  //console.log(request)

  const b = await request.json();

  // user logged in and username in entry
  if (!locals.loggedIn || locals?.user !== b.user)
    throw error(403, 'Not allowed');

  const db = locals.db;
  const c = await db.collection('sessions');
  const r = await c.deleteOne({ _id: new ObjectId(b._id), user: b.user });
  if (!r?.deletedCount) throw error(400, 'Error');

  return json(r);
}
