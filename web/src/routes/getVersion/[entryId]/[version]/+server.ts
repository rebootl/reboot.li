import { json, error } from '@sveltejs/kit';

export async function GET({ request, params, locals }) {

  const db = locals.db;

  const entryId = params.entryId;
  const version = parseInt(params.version);
  const user = 'rebootl';
  // for now versions only for logged in + own entries
  if (!locals.loggedIn || locals?.user !== user)
    throw error(403);

  const c = await db.collection('entries');

  const q = { user: user, id: entryId, deleted: false, version: version };
  const r = await c.findOne(q);
  if (!r) throw error(404, 'Not found');

  return json(r);
}
