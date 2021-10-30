import fs from 'fs';
import { unlink, rmdir } from 'fs/promises';

import path from 'path';
import crypto from 'crypto';
import { STATICDIR, MEDIADIR } from './config.js';

// setup image storage
export function storeImage(i, username) {
  return new Promise((res, rej) => {
    const randomDirName = crypto.randomBytes(20).toString('hex');
    //console.log(i);
    const imagepath = path.join(STATICDIR, MEDIADIR, username, randomDirName,
      i.name);
    console.log('saving image: ', imagepath);
    i.mv(imagepath, (err) => {
        if (err)
            rej(err);
        res({
            originalname: i.name,
            path: '/' + imagepath.replace(STATICDIR, ''),
            size: i.size
        });
    });
  });
}

export async function deleteImage(filepath) {
  const fp = path.join(STATICDIR, filepath);
  try {
    await unlink(fp);
    console.log('successfully deleted');
  } catch (error) {
    console.error('there was an error:', error.message);
    return false;
  }
  try {
    await rmdir(path.dirname(fp));
    console.log('successfully rmdir');
  } catch (error) {
    console.error('there was an error:', error.message);
    return false;
  }
  return true;
}

export function handleUpdateImages(newImages, oldImages) {
    // compare new/old ids, delete removed images
    const newIds = newImages.map(e => e.filename);
    const oldIds = oldImages.map(e => e.filename);
    for (const oldId of oldIds) {
        if (!newIds.includes(oldId)) {
            const r = oldImages.find(i => i.filename === oldId);
            deleteImage(r);
        }
    }
}
