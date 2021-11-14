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
	import SideNav from '$lib/SideNav.svelte';
  import Main from '$lib/Main.svelte';
  import { getFilteredEntries } from '$lib/filterSort';

	export let entries = [];

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

<SideNav {entries} on:change={e => navChange(e.detail)} />
<Main entries={filteredEntries} type="task"
      on:created={e => created(e.detail)} />

<style>
</style>
