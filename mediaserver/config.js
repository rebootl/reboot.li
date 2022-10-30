// settings

// server setup
const PORT = 3005;
const BASEURL = `http://localhost:${PORT}/`
const SECRET = 'secretsecret';

// auth
const COOKIENAME = 'better-reflectapp-session';

// database setup
const DBUSER = 'reboot-li-admin';
const DBPASSWORD = 'example123';
const DBURL = `mongodb://${DBUSER}:${DBPASSWORD}@localhost:27017`;
const DBNAME = 'reboot-li';

// files directory
const MEDIADIR = 'media/'

export { PORT, BASEURL, SECRET, COOKIENAME, DBUSER, DBPASSWORD, DBURL, DBNAME, MEDIADIR};
