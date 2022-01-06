export async function get(request) {

  const db = request.locals.db;

  const entryId = request.params.entryId;
  const version = parseInt(request.params.version);
  const user = 'rebootl';
  // for now versions only for logged in + own entries
  if (!request.locals.loggedIn || request.locals?.user !== user)
    return { status: 403 };

  const c = await db.collection('entries');

  const q = { user: user, id: entryId, deleted: false, version: version };
  const r = await c.findOne(q);
  if (!r) return { status: 404 };

  return {
    body: r
  };
}
