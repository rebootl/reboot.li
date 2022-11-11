export async function load({ fetch }) {

  const url = '/restore-delete';

  const res = await fetch(url);
  if (res.ok) {
    const entries = await res.json();
    return {
      entries: entries,
    };
  }

  return {
    status: res.status,
    error: new Error(`Could not load ${url}`)
  };
}