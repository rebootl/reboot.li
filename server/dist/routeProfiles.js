import express from 'express';
import { getValidProfiles } from './modelUsers.js';
const router = express.Router();
router.get('/', async (req, res) => {
    const s = parseInt(req.query.skip) || 0;
    const l = parseInt(req.query.limit) || 10;
    const r = await getValidProfiles(req.app.locals.db, s, l);
    //const c = await req.app.locals.db.collection('profiles');
    //const r = await c.find({}).skip(s).limit(l).toArray();
    return res.send({ success: true, result: r });
});
export default router;
