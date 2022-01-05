
<script>
  import Entries from '$lib/Entries.svelte';

  import { getFilteredEntries } from '$lib/filterSort';

	export let entries = [];
  export let selectedTopics = [];
  export let selectedTags = [];

  let filteredEntries = [];

  $: filterEntries(entries, selectedTopics, selectedTags);

  function filterEntries() {
    filteredEntries = getFilteredEntries(entries, selectedTopics, selectedTags);
  }
</script>

<Entries entries={filteredEntries} />
