
export async function load({ page, fetch, parent, context }) {

	const url = '/entries/news';
	//const data = await parent();

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