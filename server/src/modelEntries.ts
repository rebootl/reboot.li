
async function getPublicEntries(db, user) {
  const c = await db.collection('entries');
  const q = [
    { $match: { $and: [
      { user: user, },
      { private: false }
    ] }},
    { $sort: { pinned: -1, date: -1 }}
  ];
  return await c.aggregate(q).toArray();
}

async function getAllEntries(db, user) {
  const c = await db.collection('entries');
  const q = [
    { $match: { user: user }},
    { $sort: { pinned: -1, date: -1 }}
  ];
  return await c.aggregate(q).toArray();
}

async function getEntry(db, user, id) {
  const c = await db.collection('entries');
  return await c.findOne({ $and: [
    { user: user },
    { id: id }
  ]});
}

export { getPublicEntries, getAllEntries, getEntry };
