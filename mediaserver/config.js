// settings

// server setup
const PORT = 3005;
const BASEURL = `http://localhost:${PORT}/`

// auth
const COOKIENAME = 'better-reflectapp-session';

// database setup
const DBUSER = 'better-reflectapp-admin';
const DBPASSWORD = 'example123';
const DBURL = `mongodb://${DBUSER}:${DBPASSWORD}@localhost:27017`;
const DBNAME = 'better-reflectapp';

// files directory
const MEDIADIR = 'media/'

export { PORT, BASEURL, COOKIENAME, DBUSER, DBPASSWORD, DBURL, DBNAME, MEDIADIR};
