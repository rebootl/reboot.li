import { fail, redirect } from '@sveltejs/kit';
// import prisma from '$lib/server/prisma';
import bcrypt from 'bcrypt';
import { v4 as uuidv4 } from 'uuid';

import { COOKIENAME } from '$env/static/private';
import { getUser } from '$lib/server/db.js';

/** @type {import('./$types').PageServerLoad} */
export async function load({ locals }) {
	if (locals.user) throw redirect(303, '/');
}

/** @type {import('./$types').Actions} */
export const actions = {
	default: async ({ cookies, request }) => {
		const data = await request.formData();

    const username = data.get('username');
    const password = data.get('password');

    if (!username || !password) return fail(401, { username, missing: true });

    const r = getUser(username);
    if (!r) return fail(401, { username, missing: true });

    const check = await bcrypt.compare(password as string, r.pwhash);

    if (check) {
      console.log("login ok");

      // uuid
      const uuid = uuidv4();
      // session -> db
      const c = await prisma.session.create({
        data: {
          uuid: uuid,
            userId: r.id,
            userAgent: request.headers.get('user-agent') ?? '',
            host: (request.headers.get('x-forwarded-for') || request.headers.get('origin')) ?? '',
        }
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
    return fail(401, { username, invalid: true });
	}
};
