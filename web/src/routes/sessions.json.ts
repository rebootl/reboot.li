
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
