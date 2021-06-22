import express from 'express';
import session from 'express-session';
import connectMongo from 'connect-mongo';
const MongoStore = connectMongo(session);
import compression from 'compression';
import fileupload from 'express-fileupload';

//import cors from 'cors';

// routes
//import routeUrlInfo from './routeUrlInfo.js'
//import routeUploadMultiImages from './routeUploadMultiImages.js';

// import auth from './auth.js';
import profiles from './routeProfiles.js';
import entries from './routeEntries.js';
// import menu from './menu.js';

// db
import { client } from './db.js';
import getDb from './db.js';

// config
import * as config from '../config.js';

// setup app

const app = express();

/*app.use(cors({
  origin: function(origin, callback){callback(null, true)},
  credentials: true
}))*/
app.use(express.json({limit: '10mb'}));
app.use(session({
  secret: config.SECRET,
  resave: false,
  saveUninitialized: false,
  cookie: {
    //maxAge: 1000 * 3600 * 24 * 365 * 5,
    HttpOnly: true,
    sameSite: true,
    secure: config.PRODUCTION
  },
  store: new MongoStore({
    client: client,
    touchAfter: 24 * 3600,
    dbName: config.DBNAME,
    collection: 'sessions'
  })
}));
app.use(compression());
app.use(fileupload({
  createParentPath: true
}));

// static files (incl. client)
app.use('/', express.static(config.STATICDIR,));

// routes w/o db access

//app.get('/api/urlinfo', async (req: any, res: any) => routeUrlInfo(req, res));
//app.post('/api/uploadMultiImages',
//  async (req : any, res : any) => routeUploadMultiImages(req, res));

// routes that need db access

async function main() {
  const db = await getDb();
  app.locals.db = db;

  // set mongodb indexes for query performance
  // (from pipolog)
  /*const l = await db.collection('logdata');
  l.createIndexes([{ key: {
    agent_timestamp: 1,
    host: 1,
    name: 1
  }}]);*/

  // app.use('/api/auth', auth);
  app.use('/api/profiles', profiles);
  app.use('/api/entries', entries);
  // app.use('/api/menu', menu);

  app.listen(config.PORT);
  console.log("Listening on port: " + config.PORT);
}
main();
