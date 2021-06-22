// -> prolly do not use this
async function getUser(db, username) {
    const c = await db.collection('users');
    return await c.findOne({ name: username });
}
async function getValidUser(db, username) {
    const c = await db.collection('users');
    const u = await c.findOne({ $and: [
            { username: username },
            { active: true }
        ] });
    // check user validity
    //if (!u.active) return null;
    return u;
}
async function getValidProfiles(db, skip, limit) {
    const c = await db.collection('users');
    const r = await c.find({ active: true }, { projection: { username: 1, admin: 1, profile: 1 } }).skip(skip).limit(limit).toArray();
    return r;
}
export { getUser, getValidUser, getValidProfiles };
