export async function load({ fetch }) {

  const url = '/sessions';

  const res = await fetch(url);
  if (res.ok) {
    return {
      sessions: await res.json(),
    };
  }

  return {
    status: res.status,
    error: new Error(`Could not load ${url}`)
  };
}