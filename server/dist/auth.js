import express from 'express';
import bcrypt from 'bcrypt';
import * as config from '../config.js';
const router = express.Router();
router.post('/login', (req, res) => {
    const username = req.body.username;
    const password = req.body.password;
    if (!username || !password)
        return res.sendStatus(401);
    if (username !== config.USER.username)
        return res.sendStatus(401);
    bcrypt.compare(password, config.USER.pwhash).then((check) => {
        if (check) {
            console.log("login ok");
            req.session.loggedIn = true;
            req.session.username = username;
            res.send({ success: true });
        }
        else {
            console.log("login failed");
            res.sendStatus(401);
        }
    });
});
router.get('/logout', (req, res) => {
    console.log(req.session);
    if (!req.session.loggedIn) {
        console.log('not logged in');
        return res.sendStatus(401);
    }
    req.session.destroy((err) => {
        if (err) {
            console.log('error destroying session: ', err);
            return res.sendStatus(500);
        }
        res.clearCookie('connect.sid', {
            httpOnly: true,
            sameSite: true
        }).send({ success: true });
        console.log('session destroyed');
    });
});
export default router;
