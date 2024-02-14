import { error, fail, redirect } from '@sveltejs/kit';
import path from 'path';

import { getEntry, createEntryDB, updateEntryDB, deleteEntryDB,
  insertImagesDB, getImageDB, deleteImageDB, updateImageCommentDB} from '$lib/server/db.js';
import { storeImage, deleteImage } from '$lib/server/imageStorage.js';

/** @type {import('./$types').PageServerLoad} */
export async function load({ locals, params }) {

  // console.log('locals', locals);
  // console.log('params', params);

  if (!locals.user) throw error(401, 'Unauthorized');

  if (params.id === 'new') {
    return { entry: null };
  }

  const entryId = parseInt(params.id) ?? 0;
  const r = getEntry(locals.user.id, entryId, true);
  if (!r) throw error(404, 'Not found');

  return { entry: r };
}

/** handle images
  * @param {File[]} images
  * @param {string[]} imageComments
  * @param {string[]} imagesExifData
  * @param {number|bigint} entryId
  * @param {number} userId
  * @param {string} username
  * @returns {Promise<boolean>} success
  */
async function handleImages(images, imageComments, imagesExifData, entryId, userId, username) {

  // store images on fs
  let rs = [];
  try {
    rs = await Promise.all(images.map(async (i) => await storeImage(i, username)));
    // console.log('rs', rs);
  } catch (error) {
    console.error(error);
    throw error;
  }

  // insert images into db
  const imageData = rs.map((r, i) => ({
    path: r.url,
    comment: String(imageComments[i] ?? ''),
    previewData: r.previewData,
    exifData: String(imagesExifData[i] ?? ''),
  }));
  const r2 = insertImagesDB(imageData, entryId, userId);
  if (!r2) throw new Error('Error inserting images into db');

  return r2;
}

/** @type {import('./$types').Actions} */
export const actions = {
  async createEntry({ locals, request }) {
    if (!locals.user) throw error(401, 'Unauthorized');

    const data = await request.formData();
    
    const manualDate = data.get('manualdate') ? new Date(String(data.get('manualdate'))) : undefined;

    const content = data.get('content');
    const comment = data.get('comment');
    
    const isPrivate = data.get('isPrivate') ? 1 : 0;

    const r = createEntryDB({
      userId: locals.user.id,
      type: 'event',
      title: '',
      content: typeof content === 'string' ? content : '',
      comment: typeof comment === 'string' ? comment : '',
      private: isPrivate,
      pinned: 0,
      manualDate,
      tags: [],
    });

    if (r.changes === 0) throw error(400, 'Error creating entry');
    
    // handle images
    const images = /** @type {File[]} */ (data.getAll('images') ?? []);
    const imageComments = /** @type {string[]} */ (data.getAll('imagecomment') ?? []);
    const imagesExifData = /** @type {string[]} */ (data.getAll('imageexifdata') ?? []);

    try {
      await handleImages(images, imageComments, imagesExifData, r.lastInsertRowid, locals.user.id, locals.user.name);
    } catch (err) {
      console.error(err);
      throw error(400, 'Something went wrong while handling images');
    }

    console.log('r', r);
    redirect(303, '/timeline')
  },
  async updateEntry({ locals, request, params }) {
    if (!locals.user) throw error(401, 'Unauthorized');

    const data = await request.formData();
    
    const manualDate = data.get('manualdate') ? new Date(String(data.get('manualdate'))) : undefined;

    const content = data.get('content');
    const comment = data.get('comment');

    const isPrivate = data.get('isPrivate') ? 1 : 0;

    const entryId = parseInt(params.id) ?? 0;

    const r = updateEntryDB({
      entryId,
      userId: locals.user.id,
      type: 'event',
      title: '',
      content: typeof content === 'string' ? content : '',
      comment: typeof comment === 'string' ? comment : '',
      private: isPrivate,
      pinned: 0,
      manualDate,
      tags: [],
    });

    // console.log('r', r);
    if (r.changes === 0) throw error(400, 'Error updating entry');

    // handle images
    const images = /** @type {File[]} */ (data.getAll('images') ?? []);
    const imageComments = /** @type {string[]} */ (data.getAll('imagecomment') ?? []);

    try {
      await handleImages(images, imageComments, entryId, locals.user.id, locals.user.name);
    } catch (err) {
      console.error(err);
      throw error(400, 'Something went wrong while handling images');
    }

    redirect(303, `/timeline`)
  },
  async deleteEntry({ locals, params }) {
  
    if (!locals.user) throw error(401, 'Unauthorized');
    
    // delete images from fs
    const entry = getEntry(locals.user.id, parseInt(params.id) ?? 0, true);
    if (!entry) throw error(404, 'Not found');

    const images = entry.images ?? [];
    for (const i of images) {
      // delete image from fs
      const imagepath = path.join('static', i.path);

      const r = await deleteImage(imagepath);
      if (!r) throw error(400, `Error deleting image ${i.id} from fs`);

      // delete entry from db
      const r2 = deleteImageDB(i.id, locals.user.id);
      if (r2.changes === 0) throw error(400, `Error deleting image ${i.id} from db`);
    }

    // delete entry from db
    const r = deleteEntryDB(parseInt(params.id) ?? 0, locals.user.id);

    console.log('r', r);
    if (r.changes === 0) throw error(400, 'Error deleting entry');
    redirect(303, `/timeline`)
  },
  async deleteImage({ locals, params, request }) {
    if (!locals.user) throw error(401, 'Unauthorized');

    const data = await request.formData();
    const imageId = parseInt(String(data.get('imageId'))) ?? 0;

    // delete image from fs
    const image = getImageDB(imageId, locals.user.id);
    if (!image) throw error(404, 'Not found');

    const imagepath = path.join('static', image.path);

    const r = await deleteImage(imagepath);
    if (!r) throw error(400, `Error deleting image ${imageId} from fs`);

    // delete entry from db
    const r2 = deleteImageDB(imageId, locals.user.id);
    if (r2.changes === 0) throw error(400, `Error deleting image ${imageId} from db`);

    redirect(303, `/editEvent/${params.id}`)
  },
  async updateImageComment({ locals, params, request }) {
    if (!locals.user) throw error(401, 'Unauthorized');

    const data = await request.formData();
    const imageId = parseInt(String(data.get('imageid') ?? 0));
    const imageComment = String(data.get('imagecomment') ?? '');

    // update entry from db
    const r = updateImageCommentDB(imageId, locals.user.id, imageComment);
    if (r.changes === 0) throw error(400, `Error updating image ${imageId} in db`);

    redirect(303, `/editEvent/${params.id}`)
  },
}
