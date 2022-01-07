
<script>
  import Entries from '$lib/Entries.svelte';

  import { getFilteredEntries } from '$lib/filterSort';

	export let entries = [];
  export let selectedTopics = [];
  export let selectedTags = [];
  export let filterText = '';

  let filteredEntries = [];

  $: filterEntries(entries, selectedTopics, selectedTags, filterText);

  function filterEntries() {
    filteredEntries = getFilteredEntries(entries, selectedTopics, selectedTags,
      filterText);
  }
</script>

<Entries entries={filteredEntries} />
