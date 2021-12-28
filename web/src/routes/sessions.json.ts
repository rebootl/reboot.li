
export async function get(request) {

  if (!request.locals.loggedIn)
    return { status: 403 };

  const db = request.locals.db;

  const c = await db.collection('sessions');

  const q = { user: request.locals.user };

  const r = await c.find(q).sort({ createdAt: -1 }).toArray();
  if (!r) return { status: 400 };

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
  const r = await c.deleteOne({ uuid: b.uuid });
  if (!r?.deletedCount) return { status: 400 };

  return {
    body: {
      success: true,
      result: r
    }
  };
}
