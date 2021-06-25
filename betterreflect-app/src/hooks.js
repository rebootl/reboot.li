import { getDb } from '$lib/db';

let db;

async function initDb() {
  db = await getDb();
}
initDb();

export async function handle({ request, resolve }) {

	request.locals.db = db;

	const response = await resolve(request);

	return {
		...response,
	};
}
