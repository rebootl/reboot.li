export async function get(request) {

  if (!request.locals.loggedIn)
    return { status: 403 };

  const db = request.locals.db;

  const c = await db.collection('entries');

  const q = { deleted: true };

  const r = await c.find(q).sort({ date: -1 }).toArray();
  if (!r) return { status: 404 };

  return {
    body: r
  };
}

export async function put(request) {
  //console.log(request)

  const b = request.body;

  // user logged in and username in entry
  if (!request.locals.loggedIn || request.locals?.user !== b.user)
    return { status: 403 }

  const db = request.locals.db;
  const c = await db.collection('entries');
  //const r = await c.deleteOne({ id: b.id });
  const r = await c.updateOne({ id: b.id }, { $set: {
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
  const r = await c.deleteOne({ id: b.id });
  if (!r?.deletedCount) return { status: 400 };

  return {
    body: {
      success: true,
      result: r
    }
  };
}
