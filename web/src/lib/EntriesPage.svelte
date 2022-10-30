<script>
	import SideNav from '$lib/SideNav.svelte';
  import NewEntry from '$lib/NewEntry.svelte';
  import EntriesFilter from '$lib/EntriesFilter.svelte';
  import { refs } from '$lib/refs';
  //import { session } from '$app/stores';
  import { page } from "$app/stores";

	export let entries = [];
  export let type = 'task';

  let selectedTopics = [];
  let selectedTags = [];
	let filterText = '';

  function navChange(v) {
    selectedTopics = v[0];
    selectedTags = v[1];
		filterText = v[2];
  }

  function created(e) {
    entries.push(e);
    entries = entries;
  }
</script>

<SideNav {entries} on:change={e => navChange(e.detail)} />
<main class="margin-left">
	<h1>{refs[type].text}</h1>
  {#if $page.data.loggedIn}
    <NewEntry {type} on:created={e => created(e.detail)} />
  {/if}
  <EntriesFilter {entries} {selectedTopics} {selectedTags} {filterText} />
</main>
