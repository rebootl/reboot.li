// settings

const PRODUCTION = false;

// server port
const PORT = 4041;

// random secret
// -> how to create best?
const SECRET = 'abcdefgh';

const USER = {
  name: 'rebootl',
  pwhash: '$2b$10$wbm.5m27QVoQKvVh1Lar4uabKplVvoZFGjKuKYFCQfqilkZ5ij9oi'
};
// create with:
// const bcrypt = require('bcrypt');
// bcrypt.hashSync('beboop', 10);

// database setup
const DBUSER = 'reflectapp-svelte-admin';
const DBPASSWORD = 'example123';
const DBURL = 'mongodb://localhost:27017';
const DBNAME = 'reflectapp';

// files / paths
const STATICDIR = 'client/public';
// (mediaDir is below staticDir)
const MEDIADIR = 'media';

export { PRODUCTION, PORT, SECRET, USER, DBUSER, DBPASSWORD, DBURL,
  DBNAME, STATICDIR, MEDIADIR  };
