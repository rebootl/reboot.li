import express from 'express';
import compression from 'compression';
import fileupload from 'express-fileupload';
import cookie from 'cookie';
import cors from 'cors';

import getDb from './db.js';
import { storeImage } from './imageStorage.js';

import * as config from './config.js';

const app = express();

app.use(cors({
  origin: function(origin, callback){callback(null, true)},
  credentials: true
}));
app.use(express.json({limit: '10mb'}));

app.use(compression());
app.use(fileupload({
  createParentPath: true
}));

async function checkLogin(request, db) {
  const cookies = cookie.parse(request.headers.cookie || '');
  console.log(cookies)
  let r = null;
  if (cookies.hasOwnProperty(config.COOKIENAME) && db) {
    const c = await db.collection('sessions');
    r = await c.findOne({
      uuid: cookies[config.COOKIENAME]
    });
  }
  if (r) {
    return {
      loggedIn: true,
      user: r.user,
      admin: r.admin,
      sessionId: cookies[config.COOKIENAME]
    }
  } else {
    return {
      loggedIn: false
    }
  }
}

app.post('/api/uploadImages', async (req, res) => {
  //console.log(req)

  // check cookie
  const l = await checkLogin(req, app.locals.db);
  if (!l.loggedIn)
    return res.sendStatus(403);

  if (!req.files || Object.keys(req.files).length === 0) {
      //console.log(req.files);
      return res.status(400).send('No files were uploaded.');
  }
  let files = [];
  const filedata = req.files.filedata;
  if (Array.isArray(filedata)) {
      files = await Promise.all(filedata.map(async (f) => await storeImage(f, l.user)));
  } else {
      files.push(await storeImage(filedata, l.user));
  }
  console.log(files)
  res.send({
      success: true,
      files: files
  });
});

async function main() {
  const db = await getDb();
  app.locals.db = db;

  app.listen(config.PORT);
  console.log("Listening on port: " + config.PORT);
}
main();
