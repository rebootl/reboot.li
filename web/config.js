// settings

// auth
const COOKIENAME = 'reboot-li-session';

// for reference only
/*const USER = {
  name: 'rebootl',
  pwhash: '$2b$10$wbm.5m27QVoQKvVh1Lar4uabKplVvoZFGjKuKYFCQfqilkZ5ij9oi'
};*/
// create with:
// const bcrypt = require('bcrypt');
// bcrypt.hashSync('beboop', 10);

// database setup
const DBUSER = 'reboot-li-admin';
const DBPASSWORD = 'example123';
const DBURL = `mongodb://${DBUSER}:${DBPASSWORD}@localhost:27017`;
const DBNAME = 'reboot-li';

// mediaserver
const MEDIASERVER = 'http://localhost:3005/'
const SECRET = 'secretsecret';

export { COOKIENAME, DBUSER, DBPASSWORD, DBURL, DBNAME, MEDIASERVER, SECRET };
