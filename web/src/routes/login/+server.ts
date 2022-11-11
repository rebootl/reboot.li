import { json, error } from '@sveltejs/kit';
import bcrypt from 'bcrypt';
import { v4 as uuidv4 } from 'uuid';
import { COOKIENAME } from '$env/static/private';

export async function POST({ request, locals, cookies }) {

  const body = await request.json();

  const username = body.user;
  const password = body.password;

  if (!username || !password) throw error(401, 'login failed');

  const db = locals.db;
  const c = await db.collection('users');
  const r = await c.findOne(
    { username: username, active: true }
  );
  if (!r) throw error(401, 'login failed');

  const check = await bcrypt.compare(password, r.pwhash);

  if (check) {
    console.log("login ok");

    // uuid
    const uuid = uuidv4();
    // session -> db
    const c = await db.collection('sessions');
    const res = await c.insertOne({
      uuid: uuid,
      user: username,
      admin: r.admin,
      createdAt: new Date(),
      userAgent: request.headers.get('user-agent'),
      host: request.headers.get('x-forwarded-for') || request.headers.get('origin'),
    });

    cookies.set(
      COOKIENAME,
      uuid,
      {
        httpOnly: true,
        sameSite: 'lax',
        path: '/',
        maxAge: 60 * 60 * 24 * 365 * 10, // 10 years
        secure: true,
      }
    );

    return json(username);
  } else {
    console.log("login failed");
    throw error(401, 'login failed');;
  }
}
