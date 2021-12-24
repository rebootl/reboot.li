import bcrypt from 'bcrypt';
import cookie from 'cookie';
import { v4 as uuidv4 } from 'uuid';
import { COOKIENAME } from '../../config.js';

export async function post(request) {

  const body = JSON.parse(request.body);

  const username = body.user;
  const password = body.password;

  const error = {
    status: 401,
    body: {
      message: 'login failed'
    }
  };

  if (!username || !password) return error;

  const db = request.locals.db;
  const c = await db.collection('users');
  const r = await c.findOne(
    { username: username, active: true }
  );
  if (!r) return error;

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
      admin: r.admin
    });

    return {
      body: username,
      headers: {
        'set-cookie': [
          cookie.serialize(
            COOKIENAME,
            uuid,
            {
              httpOnly: true,
              sameSite: 'Lax',
              secure: true,
              //path: '/',
              //maxAge: 60 * 60 * 24 * 7 // 1 week
            }
          )
        ]
      }
    };
  } else {
    console.log("login failed");
    return error;
  }

}
