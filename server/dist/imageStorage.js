import fs from 'fs';
import path from 'path';
import crypto from 'crypto';
import { STATICDIR, MEDIADIR } from '../config.js';
// setup image storage
export function storeImage(i) {
    return new Promise((res, rej) => {
        const randomDirName = crypto.randomBytes(20).toString('hex');
        const imagepath = path.join(STATICDIR, MEDIADIR, randomDirName, i.name);
        console.log('saving image: ', imagepath);
        i.mv(imagepath, (err) => {
            if (err)
                rej(err);
            res({
                originalname: i.name,
                path: imagepath.replace(STATICDIR, ''),
                size: i.size
            });
        });
    });
}
export function deleteImage(image) {
    if (!image.filepath) {
        console.log("image has no filepath argument, returning");
        return image;
    }
    const fp = path.join(STATICDIR, image.filepath);
    fs.unlink(fp, (err) => {
        if (err)
            console.log('error deleting image:', err);
        console.log('deleted image', fp);
        fs.rmdir(path.dirname(fp), (err) => {
            if (err)
                console.log('error removing directory: ', err);
            console.log('directory removed:', path.dirname(fp));
        });
    });
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
