export async function get(request) {

  if (!request.locals.loggedIn || !request.locals?.user) {
    return { status: 403 };
  }

  const db = request.locals.db;

  const c = await db.collection('entries');

  const q = { user: request.locals.user, last: true, deleted: true };

  const r = await c.find(q).sort({ deleteDate: -1 }).toArray();
  if (!r) return { status: 404 };

  return {
    body: r
  };
}

export async function put(request) {
  //console.log(request)

  const b = request.body;

  // user logged in and username in entry
  if (!request.locals.loggedIn || request.locals?.user !== b.user) {
    return { status: 403 }
  }

  const db = request.locals.db;
  const c = await db.collection('entries');
  //const r = await c.deleteOne({ id: b.id });
  const r = await c.updateMany({ id: b.id, user: b.user }, { $set: {
    deleted: false
  }});
  if (!r?.modifiedCount) return { status: 400 };

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
  const r = await c.deleteMany({ user: b.user, id: b.id });
  if (!r?.deletedCount) return { status: 400 };

  return {
    body: {
      success: true,
      result: r
    }
  };
}
