import { default as mongodb } from 'mongodb';
import * as config from '../config.js';

// db setup

const MongoClient = mongodb.MongoClient;
export const client = new MongoClient(config.DBURL, {
  auth: { user: config.DBUSER, password: config.DBPASSWORD },
  useUnifiedTopology: true
});

async function getDb() {
  try {
    await client.connect();
    console.log("Connected successfully to server");
    return await client.db(config.DBNAME);
  } catch(e) {
    throw e;
  }
}

export default getDb;
