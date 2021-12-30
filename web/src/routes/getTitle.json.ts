import jsdom from 'jsdom';
const { JSDOM } = jsdom;

export async function post(request) {
  console.log('gettitle post')
  if (!request.locals.loggedIn || !request.locals.user)
    return { status: 403 };

  const b = request.body;

  try {
    const r = await fetch(b.url);
    if (r.ok) {
      const dom = new JSDOM(await r.text());
      const title = dom.window.document.querySelector('title').textContent;
      return {
        body: {
          title: title
        }
      }
    } else {
      new Error(r);
    }
  } catch(err) {
    //console.log(err)
    return {
      body: {
        title: 'error getting title',
        error: true,
        code: err.code
      }
    }
  }
}
