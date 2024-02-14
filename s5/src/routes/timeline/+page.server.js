import dayjs from 'dayjs';

import { error } from '@sveltejs/kit';
import { getEntries } from '$lib/server/db.js';

/** @type {import('./$types').PageServerLoad} */
export async function load({ locals }) {

  const type = 'event';
  const userId = 1;

  let loggedIn = false;
  if (locals.user && locals?.user.id == userId) {
    loggedIn = true;
  }

  const entries = getEntries(userId, type, loggedIn);
  if (!entries) throw error(404, 'Not found');
  if (entries.length === 0) return { timelineEntries: [] };

  let previousYear = dayjs(entries[0].created_at).format('YYYY');
  let previousMonth = dayjs(entries[0].created_at).format('MMMM');

  // we want to display a timeline with entries grouped by year and month
  // we do this by creating a new array with entries and year/month headers
  // we also sort the entries by date

  /** @type {import('$lib/types').TimelineEntry[]} */
  const timelineEntries = [];

  const sortedEntries = entries.sort((a, b) => {
    const aDate = a.manual_date || a.created_at;
    const bDate = b.manual_date || b.created_at;
    return dayjs(aDate).isBefore(dayjs(bDate)) ? 1 : -1;
  });

  for (const entry of sortedEntries) {
    const entryDate = entry.manual_date || entry.created_at;

    const entryYear = dayjs(entryDate).format('YYYY');
    const entryMonth = dayjs(entryDate).format('MMMM');
    // console.log(entryYear);
    // console.log(entryMonth);
    if (entryYear !== previousYear) {
      timelineEntries.push({
        type: 'year',
        year: entryYear,
        month: null,
        entry: null,
        date: null,
      });
      previousYear = entryYear;
    }
    if (entryMonth !== previousMonth) {
      timelineEntries.push({
        type: 'month',
        year: null,
        month: entryMonth,
        entry: null,
        date: null,
      });
      previousMonth = entryMonth;
    }

    // pre-process image exif data
    for (const image of entry.images) {
      // console.log(image.exif_data);
      if (image.exif_data) {
        // console.log(image.exif_data);
        const exifData = JSON.parse(image.exif_data);
        // convert exposure time from floating point to 1/x format
        let exposureTime = exifData?.ExposureTime;
        if (exposureTime) {
          const exposureTimeFloat = parseFloat(exposureTime);
          if (exposureTimeFloat < 1) {
            exposureTime = '1/' + Math.round(1 / exposureTimeFloat);
          }
        }
        const exifDataOut = {
          make: exifData?.Make,
          model: exifData?.Model,
          createDate: exifData?.CreateDate,
          exposureTime: exposureTime,
          fNumber: exifData?.FNumber,
          focalLength: exifData?.FocalLength,
          iso: exifData?.ISO,
          lensModel: exifData?.LensModel,
        };
        // console.log(exifDataOut);
        image.exifData = exifDataOut;
        delete image.exif_data;
      }
    }
    timelineEntries.push({
      type: 'entry',
      year: null,
      month: null,
      entry: entry,
      date: entryDate,
    });
  }

  return { timelineEntries };
}
