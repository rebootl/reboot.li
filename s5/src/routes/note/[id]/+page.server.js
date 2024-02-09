import { error } from '@sveltejs/kit';
import { getEntry } from '$lib/server/db.js';

/** @type {import('./$types').PageServerLoad} */
export async function load({ locals, params }) {

  const userId = locals.user?.id ?? 1;
  const entryId = parseInt(params.id) ?? 0;
  const loggedIn = locals.user ? true : false;

  console.log('userId:', userId);
  
  const r = getEntry(userId, entryId, loggedIn);
  if (!r) throw error(404, 'Not found');

  return { entry: r };
}
