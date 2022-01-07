import jwt from 'jsonwebtoken';
import { SECRET } from '../../config.js';

function createToken(user) : string {
  // sign with default (HMAC SHA256)
  const expirationDate =  Math.floor(Date.now() / 1000) + 30 // 30 seconds from now
  const token = jwt.sign({
    user: user,
    exp: expirationDate
  }, SECRET);
  return token;
}

export async function get(request) {

  if (!request.locals.loggedIn || !request.locals?.user) return { status: 403 };

  const t = createToken(request.locals.user);

  return {
    body: {
      token: t
    }
  }
}
