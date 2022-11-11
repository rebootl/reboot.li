import { json, error } from '@sveltejs/kit';

export async function GET({ request, params, locals }) {

  if (!locals.loggedIn || !locals?.user) {
    throw error(403, 'Not allowed');
  }

  const db = locals.db;

  const c = await db.collection('entries');

  const q = { user: locals.user, last: true, deleted: true };

  const r = await c.find(q).sort({ deleteDate: -1 }).toArray();
  if (!r) throw error(404, 'Not found');

  return json(r);
}

export async function PUT({ request, locals }) {
  //console.log(request)

  const b = await request.json();

  // user logged in and username in entry
  if (!locals.loggedIn || locals?.user !== b.user) {
    throw error(403, 'Not allowed');
  }

  const db = locals.db;
  const c = await db.collection('entries');
  //const r = await c.deleteOne({ id: b.id });
  const r = await c.updateMany({ id: b.id, user: b.user }, { $set: {
    deleted: false
  }});
  if (!r?.modifiedCount) throw error(400, 'Error');

  return json(r);
}

export async function DELETE({ request, locals }) {
  //console.log(request)

  const b = await request.json();

  // user logged in and username in entry
  if (!locals.loggedIn || locals?.user !== b.user)
    return { status: 403 }

  const db = locals.db;
  const c = await db.collection('entries');
  const r = await c.deleteMany({ user: b.user, id: b.id });
  if (!r?.deletedCount) throw error(400, 'Error');

  return json(r);
}
