import { json, error } from '@sveltejs/kit';
/*import jsdom from 'jsdom';
const { JSDOM } = jsdom;
*/

/** @type {import('@sveltejs/kit').RequestHandler} */
export async function POST({ request, locals }) {
  //console.log('gettitle post')
  if (!locals.user) return { status: 403 };

  const b = await request.json();

  try {
    const r = await fetch(b.url);
    if (r.ok) {
      //const dom = new JSDOM(await r.text());
      //const title = dom.window.document.querySelector('title').textContent;
      const text = await r.text();
      const m = text.match(/<title>(.*?)<\/title>/);
      return json({
        title: m[1] || '',
      });
    } else {
      return json({
        title: 'error getting title',
        error: true,
        message: `${r.status} ${r.statusText}`,
      });
    }
  } catch(err) {
    console.log(err);
    return json({
      title: 'error getting title',
      error: true,
      message: err.message,
    });
  }
}
