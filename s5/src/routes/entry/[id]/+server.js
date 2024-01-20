import { json, error } from '@sveltejs/kit';
import { deleteEntryDB } from '$lib/server/db.js';

/** @type {import('@sveltejs/kit').RequestHandler} */
export async function DELETE({ params, locals }) {
  //console.log(request)

  if (!locals.user) throw error(401, 'Unauthorized');

  const r = deleteEntryDB({
    entryId: parseInt(params.id) ?? 0,
    userId: locals.user.id,
  });
  
  console.log('r', r);
  if (r.changes === 0) throw error(400, 'Error deleting entry');

  return json({ result: r });
}
