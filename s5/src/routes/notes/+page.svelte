<script>
	export let data;
  console.log(data);
  
	let entries = data.entries;
</script>

{#each entries as entry}
  <a href="/entry/{entry.id}">{entry.title}</a>
{:else}
  <p>No entries yet.</p>
{/each}
