<script context="module">
  export async function load({ page, fetch, session, context }) {

  	const url = '/entries/task.json';

  	const res = await fetch(url);
  	if (res.ok) {
  		return {
  			props: {
  				entries: await res.json(),
  			}
  		};
  	}

  	return {
  		status: res.status,
  		error: new Error(`Could not load ${url}`)
  	};
  }
</script>

<script>
  import EntriesPage from '$lib/EntriesPage.svelte';

	export let entries = [];
</script>

<EntriesPage {entries} type={'task'} />
