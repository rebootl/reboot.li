import { json, error } from '@sveltejs/kit';
import jwt from 'jsonwebtoken';
import { SECRET } from '$env/static/private';

function createToken(user) : string {
  // sign with default (HMAC SHA256)
  const expirationDate =  Math.floor(Date.now() / 1000) + 30 // 30 seconds from now
  const token = jwt.sign({
    user: user,
    exp: expirationDate
  }, SECRET);
  return token;
}

export async function GET({ request, params, locals }) {

  if (!locals.loggedIn || !locals?.user) throw error(403);

  const t = createToken(locals.user);

  return json({ token: t });
}
