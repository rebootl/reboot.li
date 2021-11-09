<script context="module">
  export async function load({ page, fetch, session, context }) {

  	const url = '/entries/link.json';

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
	import Main from '$lib/Main.svelte';
  import { getFilteredEntries } from '$lib/filterSort.ts';

	export let entries = [];
  export let showSideNav = true;

  let filteredEntries = [];

  $: filterEntries([[], []], entries);

  function filterEntries(v) {
    filteredEntries = getFilteredEntries(entries, v);
  }
</script>

<SideNav {entries} hidden={showSideNav}
         on:change={e => filterEntries(e.detail)} />
<Main entries={filteredEntries} {showSideNav} type="link" />

<style>
</style>
