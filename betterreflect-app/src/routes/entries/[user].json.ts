
export async function get(request) {

  //console.log(request)

  const db = request.locals.db;

  const user = request.params.user;

  const c = await db.collection('entries');

  const q = { user: user, private: false };
  if (request.locals.loggedIn && request.locals?.user === user)
    delete q.private;

  const r = await c.find(q).sort({ date: -1 }).toArray();

  return {
    body: r
  };

}
