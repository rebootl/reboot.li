import express from 'express';
import compression from 'compression';
import fileupload from 'express-fileupload';
import expressJwt from 'express-jwt';
import cors from 'cors';
import path from 'path';

import { storeImage, deleteImage } from './imageStorage.js';

import * as config from './config.js';

const app = express();

app.use(cors({
  origin: function(origin, callback) { callback(null, true) },
  credentials: true
}));
app.use(express.json({limit: '10mb'}));

app.use(expressJwt({
  secret: config.SECRET,
  algorithms: ['HS256'],
  requestProperty: 'token',
  credentialsRequired: false
}));

app.use(compression());
app.use(fileupload({
  createParentPath: true
}));

// serve files
app.use(path.join('/', config.MEDIADIR), express.static(config.MEDIADIR));

app.post('/api/uploadImages', async (req, res) => {

  if (!req.token) {
    console.log('access denied');
    return res.sendStatus(401);
  }
  const user = req.token.user;

  if (!req.files || Object.keys(req.files).length === 0) {
      return res.status(400).send('No files were uploaded.');
  }
  let files = [];
  const filedata = req.files.filedata;
  if (Array.isArray(filedata)) {
      files = await Promise.all(filedata.map(async (f) => await storeImage(f, user)));
  } else {
      files.push(await storeImage(filedata, user));
  }
  res.send({
      success: true,
      files: files
  });
});

app.post('/api/deleteImage', async (req, res) => {

  if (!req.token) {
    console.log('access denied');
    return res.sendStatus(401);
  }
  const user = req.token.user;

  const fp = req.body.filepath;
  const userdir = fp.split('/')[1];
  if (userdir !== user)
    return res.sendStatus(403);

  const r = await deleteImage(fp);
  if (!r)
    return res.send({
      success: false,
    });

  res.send({
    success: true,
  });
});

async function main() {
  app.listen(config.PORT);
  console.log("Listening on port: " + config.PORT);
}
main();
