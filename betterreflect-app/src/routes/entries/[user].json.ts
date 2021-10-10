
export async function get(request) {

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

const allowedTypes = [ 'task', 'link', 'article', 'image' ];
const requiredFields = [ 'id', 'date', 'user', 'type', 'topics', 'tags',
  'private', 'pinned' ];

export async function post(request) {
  console.log(request)

  const b = request.body;

  // user logged in and username in entry
  if (!request.locals.loggedIn || request.locals?.user !== b.user)
    return { status: 403 }

  // checks
  for (const f of requiredFields) {
    if (!b.hasOwnProperty(f)) return { status: 403 };
  }

  if (!allowedTypes.includes(b.type)) return { status: 403 };

  const db = request.locals.db;
  const c = await db.collection('entries');
  const r = await c.insertOne(b);
  console.log(r)

  return {
    body: r.ops[0]
  };
}
