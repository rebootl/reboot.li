import { error, redirect } from '@sveltejs/kit';
import bcrypt from 'bcrypt';
import { v4 as uuidv4 } from 'uuid';

import { COOKIENAME } from '$env/static/private';
import { getUser, createSession } from '$lib/server/db.js';

/** @type {import('./$types').PageServerLoad} */
// export async function load({ locals }) {

//   /** @typedef {{ loggedIn: boolean, username: string | null }} */
// 	// if (locals.user) throw redirect(303, '/');

//   return {};
// }

/** @type {import('./$types').Actions} */
export const actions = {
	default: async ({ cookies, request }) => {
		const data = await request.formData();

    const username = data.get('username');
    const password = data.get('password');

    if (!username || !password) throw error(401, 'Login failed, username or password missing');
    
    // console.log("username", username);

    const r = getUser(/** @type {string} */ (username));
    if (!r) throw error(401, 'Login failed.');

    const loginOk = await bcrypt.compare(/** @type {string} */ (password), r.pwhash);
    console.log("loginOk", loginOk);
    
    if (loginOk) {
      // uuid
      const uuid = uuidv4();
      // session -> db
      const c = createSession({
        uuid: uuid,
        userId: r.id,
        userAgent: request.headers.get('user-agent') ?? '',
        host: (request.headers.get('x-forwarded-for') || request.headers.get('origin')) ?? '',
      });

      cookies.set(
        COOKIENAME,
        uuid,
        {
          httpOnly: true,
          sameSite: 'lax', // default
          path: '/',
          maxAge: 60 * 60 * 24 * 365 * 10, // 10 years
          secure: true,
        }
      );

      throw redirect(303, '/');
    }
    console.log("login failed");
    throw error(401, 'Login failed.');
	}
};
