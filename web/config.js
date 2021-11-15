// settings

const PRODUCTION = false;

// server port
const PORT = 4044;

// random secret
// -> how to create best?
const SECRET = 'abcdefgh';

const COOKIENAME = 'better-reflectapp-session';

/*const USER = {
  name: 'rebootl',
  pwhash: '$2b$10$wbm.5m27QVoQKvVh1Lar4uabKplVvoZFGjKuKYFCQfqilkZ5ij9oi'
};*/
// create with:
// const bcrypt = require('bcrypt');
// bcrypt.hashSync('beboop', 10);

// database setup
const DBUSER = 'better-reflectapp-admin';
const DBPASSWORD = 'example123';
const DBURL = `mongodb://${DBUSER}:${DBPASSWORD}@localhost:27018`;
const DBNAME = 'better-reflectapp';

// files / paths
const STATICDIR = 'client/public';
// (mediaDir is below staticDir)
const MEDIADIR = 'media';

export { PRODUCTION, PORT, SECRET, COOKIENAME, DBUSER, DBPASSWORD, DBURL,
  DBNAME, STATICDIR, MEDIADIR  };
