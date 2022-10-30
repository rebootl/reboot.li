import { MongoClient } from 'mongodb';
//import * as config from '../../config';
import { DBURL, DBNAME } from '$env/static/private';

// db setup
export const client = new MongoClient(DBURL, {
  useUnifiedTopology: true
});
//auth: { user: config.DBUSER, password: config.DBPASSWORD },

export async function getDb() {
  try {
    await client.connect();
    console.log("Connected successfully to server");
    return await client.db(DBNAME);
  } catch(e) {
    throw e;
  }
}
