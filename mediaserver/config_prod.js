// settings

const PRODUCTION = true;

// server port
const PORT = 3005;

// random secret
// -> how to create best?
const SECRET = 'abcdefgh';

const COOKIENAME = 'better-reflectapp-session';


// database setup
const DBUSER = 'better-reflectapp-admin';
const DBPASSWORD = 'example123';
const DBURL = `mongodb://${DBUSER}:${DBPASSWORD}@db:27017`;
const DBNAME = 'better-reflectapp';

const STATICDIR = PRODUCTION ? 'media/' :
  '../web/static/';
const MEDIADIR = 'media'

export { PRODUCTION, PORT, SECRET, COOKIENAME, DBUSER, DBPASSWORD, DBURL,
  DBNAME, STATICDIR, MEDIADIR  };
