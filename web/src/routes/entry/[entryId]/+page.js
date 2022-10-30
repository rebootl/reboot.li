
export async function load({ url, params, fetch }) {

  console.log(url)

  const entryId = params.entryId;
  const fetchURL = `/entry/${entryId}`;
  const edit = url.searchParams.has('edit');

  console.log('load entry:')
  console.log(entryId)

  const res = await fetch(fetchURL);
  if (res.ok) {
    return {
      entry: await res.json(),
      edit: edit,
    };
  }

  return {
    status: res.status,
    error: new Error(`Could not load ${fetchURL}`)
  };
}