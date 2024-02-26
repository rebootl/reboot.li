import { error } from '@sveltejs/kit';
import { getEntries, getTagsDB } from '$lib/server/db.js';

/** @type {import('./$types').PageServerLoad} */
export async function load({ locals }) {

  const type = 'link';
  const userId = 1;

  let loggedIn = false;
  if (locals.user && locals?.user.id == userId) {
    loggedIn = true;
  }

  const r = getEntries(userId, type, loggedIn);
  if (!r) throw error(404, 'Not found');

  const tags = getTagsDB(userId);

  return {
    entries: r.reverse(),
    tags,
  };
}
