import { json, error } from '@sveltejs/kit';
import { allowedTypes } from '$lib/entryTypes.js';

export async function GET({ request, params, locals }) {

  const db = locals.db;

  const entryId = params.entryId;
  const user = 'rebootl';

  const c = await db.collection('entries');

  const q = { user: user, id: entryId, deleted: false, last: true };
  if (!locals.loggedIn || locals?.user !== user) {
    q.private = false;
  }

  const r = await c.findOne(q);
  //if (!r) return { status: 404 };
  if (!r) throw error(404, 'Not found');

  return json(r);
}

const requiredFields = [ 'id', 'date', 'user', 'type', 'topics', 'tags',
  'private', 'pinned', 'last', 'version' ];

export async function POST({ request, locals }) {
  //console.log(request)

  const b = await request.json();

  // user logged in and username in entry
  if (!locals.loggedIn || locals?.user !== b.user) {
    throw error(403);
  }

  // checks
  for (const f of requiredFields) {
    if (!b.hasOwnProperty(f)) throw error(400);
  }

  if (!allowedTypes.includes(b.type)) throw error(400);

  const db = locals.db;
  const c = await db.collection('entries');
  const r = await c.insertOne(b);
  if (!r) throw error(400);

  return json({ result: r });
}

export async function PUT({ request, locals }) {
  //console.log(request)

  const b = await request.json();

  // user logged in and username in entry
  if (!locals.loggedIn || locals?.user !== b.user) {
    throw error(403);
  }

  // checks
  for (const f of requiredFields) {
    if (!b.hasOwnProperty(f)) throw error(400);
  }

  if (!allowedTypes.includes(b.type)) throw error(400);

  // immutable field has to be deleted
  delete b._id;

  const db = locals.db;
  const c = await db.collection('entries');
  const r = await c.updateOne(
    { user: b.user, id: b.id, version: b.version },
    { $set: {
      last: false,
    }
  });
  if (!r) throw error(400);

  b.version++;

  const rn = await c.insertOne(b);
  if (!rn) throw error(400);

  return json({ result: r });
}

export async function DELETE({ request, locals }) {
  //console.log(request)

  const b = await request.json();

  // user logged in and username in entry
  if (!locals.loggedIn || locals?.user !== b.user) {
    throw error(403);
  }

  const db = locals.db;
  const c = await db.collection('entries');
  //const r = await c.deleteOne({ id: b.id });
  const r = await c.updateMany({ user: b.user, id: b.id }, { $set: {
    deleted: true,
    deleteDate: new Date()
  }});
  if (!r?.modifiedCount) throw error(400);

  return json({ result: r });
}
