import { MongoClient } from 'mongodb';
import * as config from '../../config';

// db setup
export const client = new MongoClient(config.DBURL, {
  useUnifiedTopology: true
});
//auth: { user: config.DBUSER, password: config.DBPASSWORD },

export async function getDb() {
  try {
    await client.connect();
    console.log("Connected successfully to server");
    return await client.db(config.DBNAME);
  } catch(e) {
    throw e;
  }
}
