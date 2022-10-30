
export async function load({ fetch }) {

	const url = '/entries/note';

	const res = await fetch(url);
	if (res.ok) {
		return {
			entries: await res.json(),
		};
	}

	return {
		status: res.status,
		error: new Error(`Could not load ${url}`)
	};
}