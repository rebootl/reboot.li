import fs from 'fs';
import { unlink, rmdir } from 'fs/promises';
import path from 'path';
import crypto from 'crypto';
import { MEDIADIR, BASEURL } from './config.js';

// setup image storage
export function storeImage(i, username) {
  return new Promise((res, rej) => {
    const randomDirName = crypto.randomBytes(20).toString('hex');
    const imagepath = path.join(MEDIADIR, username, randomDirName, i.name);
    console.log('saving image: ', imagepath);
    i.mv(imagepath, (err) => {
        if (err)
            rej(err);
        res({
            originalname: i.name,
            path: imagepath,
            url: new URL(imagepath, BASEURL),
            size: i.size
        });
    });
  });
}

export async function deleteImage(filepath) {
  try {
    await unlink(filepath);
    console.log('successfully deleted: ', filepath);
  } catch (error) {
    console.error('there was an error:', error.message);
    return false;
  }
  try {
    await rmdir(path.dirname(filepath));
    console.log('successfully rmdir');
  } catch (error) {
    console.error('there was an error:', error.message);
    return false;
  }
  return true;
}
