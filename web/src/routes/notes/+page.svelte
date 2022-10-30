<script>
  import EntriesPage from '$lib/EntriesPage.svelte';

	export let data;
	let entries = data.entries;
</script>

<EntriesPage {entries} type={'note'} />
