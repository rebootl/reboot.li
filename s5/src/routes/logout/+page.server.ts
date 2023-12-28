import { fail, redirect } from '@sveltejs/kit';
import { COOKIENAME } from '$env/static/private';
import { destroySession } from '$lib/server/db.js';

/** @type {import('@sveltejs/kit').Load} */
export async function load({ locals }) {
	if (!locals.user) throw redirect(307, '/login');
}

/** @type {import('./$types').Actions} */
export const actions = {
	default: async ({ cookies, locals }) => {

    const sessionId = cookies.get(COOKIENAME);
    if (!sessionId) return fail(401, { message: 'Session not found' });

		const r = await destroySession(sessionId);
		console.log('destroySession', r);
    if (!r) return fail(401, { message: 'Session not found' });

    cookies.delete(COOKIENAME, { path: '/' });

    locals.user = null;
	  throw redirect(303, '/');
  }
};
