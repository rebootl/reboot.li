import { writeFile, mkdir, unlink, rmdir } from 'fs/promises';
import path from 'path';
import crypto from 'crypto';
import sharp from 'sharp';

import { MEDIADIR } from '$env/static/private';

/** @param {File} file
  * @param {number} maxWidth
  * @param {number} maxHeight
  * @returns {Promise<string>}
  */
export async function resizeAndEncode(file, maxWidth=1920, maxHeight=1920) {
  const buffer = await sharp(await file.arrayBuffer())
    .resize(maxWidth, maxHeight, { fit: 'inside' })
    .toBuffer();

  const encoded = buffer.toString('base64');
  return encoded;
}

// setup image storage
/** @param {File} i
  * @param {string} username
  * @returns {Promise<{originalname: string, path: string, url: string, size: number, previewData: string}>}
	*/
export async function storeImage(i, username) {
  try {
    const randomDirName = crypto.randomBytes(20).toString('hex');
    const name = i.name.replace(/[^a-zA-Z0-9_\-\.]/g, '_'); 
    // -> limit max length
    const imagepath = path.join('static', MEDIADIR, username, randomDirName, name);
    
    await mkdir(path.dirname(imagepath), { recursive: true });
    await writeFile(imagepath, Buffer.from(await i.arrayBuffer()));

    return {
      originalname: i.name,
      path: imagepath,
      url: path.join(MEDIADIR, username, randomDirName, name),
      size: i.size,
      previewData: await resizeAndEncode(i, 120, 120),
    };
  } catch (error) {
    throw error;
  }
}

/** @param {string} filepath
  * @returns {Promise<boolean>}
  */
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
 
