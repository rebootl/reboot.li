import { allowedTypes } from '$lib/entryTypes.ts';

export async function get(request) {

  const db = request.locals.db;

  const entryId = request.params.entryId;
  const user = 'rebootl';

  const c = await db.collection('entries');

  const q = { user: user, id: entryId, deleted: false };
  if (!request.locals.loggedIn || !request.locals?.user === user)
    q.private = false;

  const r = await c.findOne(q);
  if (!r) return { status: 404 };

  return {
    body: r
  };
}

const requiredFields = [ 'id', 'date', 'user', 'type', 'topics', 'tags',
  'private', 'pinned' ];

export async function post(request) {
  //console.log(request)

  const b = request.body;

  // user logged in and username in entry
  if (!request.locals.loggedIn || request.locals?.user !== b.user)
    return { status: 403 }

  // checks
  for (const f of requiredFields) {
    if (!b.hasOwnProperty(f)) return { status: 400 };
  }

  if (!allowedTypes.includes(b.type)) return { status: 400 };

  const db = request.locals.db;
  const c = await db.collection('entries');
  const r = await c.insertOne(b);
  if (!r) return { status: 400 };

  return {
    body: {
      result: r
    }
  };
}

export async function put(request) {
  //console.log(request)

  const b = request.body;

  // user logged in and username in entry
  if (!request.locals.loggedIn || request.locals?.user !== b.user)
    return { status: 403 }

  // checks
  for (const f of requiredFields) {
    if (!b.hasOwnProperty(f)) return { status: 400 };
  }

  if (!allowedTypes.includes(b.type)) return { status: 400 };

  // immutable field has to be deleted
  delete b._id;

  const db = request.locals.db;
  const c = await db.collection('entries');
  const r = await c.replaceOne({ id: b.id }, b);
  if (!r) return { status: 400 };

  return {
    body: {
      result: r
    }
  };
}

export async function del(request) {
  //console.log(request)

  const b = request.body;

  // user logged in and username in entry
  if (!request.locals.loggedIn || request.locals?.user !== b.user)
    return { status: 403 }

  const db = request.locals.db;
  const c = await db.collection('entries');
  //const r = await c.deleteOne({ id: b.id });
  const r = await c.updateOne({ id: b.id }, { $set: {
    deleted: true,
    deleteDate: new Date()
  }});
  if (!r?.modifiedCount) return { status: 400 };

  return {
    body: {
      result: r
    }
  };
}
