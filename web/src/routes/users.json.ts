


export async function get(request) {

  //console.log(request)
  const db = request.locals.db;

  const c = await db.collection('users');
  const r = await c.find(
    { active: true },
    { projection: { _id: 0, username: 1, profile: 1 }}
  ).toArray();

  return {
    body: r
  };

}
