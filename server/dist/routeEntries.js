import express from 'express';
import { getPublicEntries, getAllEntries, getEntry } from './modelEntries.js';
import { getValidUser } from './modelUsers.js';
import { handleUpdateImages, deleteImage } from './imageStorage.js';
const router = express.Router();
const allowedTypes = ['any', 'task', 'link', 'article', 'image'];
const requiredFields = ['id', 'type', 'date', 'topics', 'tags'];
router.get('/:username', async (req, res) => {
    const db = req.app.locals.db;
    const user = req.params.username;
    // check if requested user exists _and valid_ !!
    const u = await getValidUser(db, user);
    if (!u)
        return res.sendStatus(404);
    //const skip = parseInt(req.query.skip) || 0;
    //const limit = parseInt(req.query.limit) || 0;
    let r;
    if (!req.session.loggedIn || req.session.username !== user) {
        // not logged in or requesting different user
        r = await getPublicEntries(db, user);
    }
    else {
        r = await getAllEntries(db, user);
    }
    return res.send({ success: true, result: r });
});
router.post('/', async (req, res) => {
    if (!req.session.loggedIn || req.session.username !== req.body.user) {
        console.log('unallowed entry post rejected');
        res.sendStatus(401);
        return;
    }
    //console.log(req.session);
    //console.log(req.body);
    const db = req.app.locals.db;
    const c = await db.collection('entries');
    // check required fields
    for (const f of requiredFields) {
        if (!req.body.hasOwnProperty(f))
            return res.sendStatus(400);
    }
    // more checks?
    // - length
    // get existing images for handleUpdateImages below
    let oldImages = [];
    const rOld = await getEntry(db, req.body.user, req.body.id);
    if (rOld)
        if (rOld.images)
            oldImages = rOld.images;
    // insert into db
    const r = await c.updateOne({ $and: [
            { user: req.body.user },
            { id: req.body.id }
        ] }, { $set: { ...req.body } }, { upsert: true });
    if (!r)
        return res.sendStatus(400);
    // handle images
    // (this deletes removed images on server,
    //  upload of new images is separate via image upload)
    let newImages = [];
    if (req.body.images)
        newImages = req.body.images;
    handleUpdateImages(newImages, oldImages);
    return res.send({ success: true, result: r });
});
router.delete('/', async (req, res) => {
    if (!req.session.loggedIn || req.session.username !== req.body.user) {
        console.log('unallowed entry delete rejected');
        res.sendStatus(401);
        return;
    }
    //console.log(req.session);
    const db = req.app.locals.db;
    const c = await db.collection('entries');
    // check required fields
    if (!req.body.hasOwnProperty('id'))
        return res.sendStatus(400);
    // more checks?
    // - length
    // delete from db
    const r = await c.remove({ id: req.body.id });
    if (!r)
        return res.sendStatus(400);
    // delete images
    for (const image of req.body.images) {
        deleteImage(image);
    }
    return res.send({ success: true, result: r });
});
export default router;
