<script context="module">
  export async function load({ page, fetch, session, context }) {

  	const url = '/entries/image.json';

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
	import SideNav from '$lib/SideNav.svelte';
	import Entries from '$lib/Entries.svelte';
  import { getFilteredEntries } from '$lib/filterSort';

	export let entries = [];
  export let showSideNav = true;

  let filteredEntries = entries;

  function filterEntries(v) {
    filteredEntries = getFilteredEntries(entries, v);
  }
</script>

<SideNav {entries} hidden={showSideNav}
         on:change={e => filterEntries(e.detail)} />
<Entries entries={filteredEntries} {showSideNav} />

<style>
</style>
