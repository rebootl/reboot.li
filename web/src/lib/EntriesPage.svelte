<script>
	import SideNav from '$lib/SideNav.svelte';
  import NewEntry from '$lib/NewEntry.svelte';
  import EntriesFilter from '$lib/EntriesFilter.svelte';

  import { session } from '$app/stores';

	export let entries = [];
  export let type = 'task';

  let selectedTopics = [];
  let selectedTags = [];

  function navChange(v) {
    selectedTopics = v[0];
    selectedTags = v[1];
  }

  function created(e) {
    entries.push(e);
    entries = entries;
  }
</script>

<SideNav {entries} on:change={e => navChange(e.detail)} />
<main class="margin-left">
  {#if $session.loggedIn}
    <NewEntry {type} on:created={e => created(e.detail)} />
  {/if}
  <EntriesFilter {entries} {selectedTopics} {selectedTags} />
</main>
