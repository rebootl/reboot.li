import { allowedTypes } from '$lib/entryTypes.ts';

export async function get(request) {

  const db = request.locals.db;

  const type = request.params.type;
  const user = 'rebootl';

  const c = await db.collection('entries');

  const q = { user: user, deleted: false, last: true };
  if (allowedTypes.includes(type))
    q.type = type;
  if (!request.locals.loggedIn || !request.locals?.user === user)
    q.private = false;

  const r = await c.find(q).sort({ date: -1 }).toArray();
  if (!r) return { status: 404 };

  return {
    body: r
  };
}
