import cookie from 'cookie';
import { COOKIENAME } from '../../config.js';

export async function post(request) {

  const error = {
    status: 401,
    body: {
      message: 'logout failed'
    }
  };

  if (!request.locals.loggedIn) {
    return error;
  }

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
              sameSite: true,
              //secure: true,
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
