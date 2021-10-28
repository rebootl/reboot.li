// settings

const PRODUCTION = false;

// server port
const PORT = 3005;

// random secret
// -> how to create best?
const SECRET = 'abcdefgh';

// database setup
const DBUSER = 'reflectapp-svelte-admin';
const DBPASSWORD = 'example123';
const DBURL = 'mongodb://localhost:27017';
const DBNAME = 'reflectapp';

const MEDIADIR = PRODUCTION ? '../betterreflect-app/build/assets/media' :
  '../betterreflect-app/static/media';

export { PRODUCTION, PORT, SECRET, DBUSER, DBPASSWORD, DBURL,
  DBNAME, MEDIADIR  };
