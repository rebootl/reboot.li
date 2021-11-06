export async function get(request) {

  const db = request.locals.db;
  console.log(request)
  const entryId = request.params.entryId;
  const user = 'rebootl';

  const c = await db.collection('entries');

  const q = { user: user, id: entryId };
  if (!request.locals.loggedIn || !request.locals?.user === user)
    q.private = false;

  const r = await c.findOne(q);
  if (!r) return { status: 404 };

  return {
    body: r
  };
}
