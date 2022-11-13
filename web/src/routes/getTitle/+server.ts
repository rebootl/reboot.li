import { json, error } from '@sveltejs/kit';
/*import jsdom from 'jsdom';
const { JSDOM } = jsdom;
*/
export async function POST({ request, locals }) {
  //console.log('gettitle post')
  if (!locals.loggedIn) return { status: 403 };

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
      const err = new Error(r.status);
      err.code = `${r.status} ${r.statusText}`;
      throw(err);
    }
  } catch(err) {
    return json({
      title: 'error getting title',
      error: true,
      code: err.code
    });
  }
}
