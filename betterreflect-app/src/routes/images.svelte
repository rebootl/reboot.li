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
  import Main from '$lib/Main.svelte';
  import { getFilteredEntries } from '$lib/filterSort';

	export let entries = [];
  export let showSideNav = true;

  let filteredEntries = [];
  let selectedTopics = [];
  let selectedTags = [];

  $: filterEntries(entries);

  function filterEntries() {
    filteredEntries = getFilteredEntries(entries, selectedTopics, selectedTags);
  }

  function navChange(v) {
    selectedTopics = v[0];
    selectedTags = v[1];
    filterEntries();
  }

  function created(e) {
    entries.push(e);
    entries = entries;
  }
</script>

<SideNav {entries} hidden={showSideNav}
         on:change={e => navChange(e.detail)} />
<Main entries={filteredEntries} {showSideNav} type="image" />

<style>
</style>
