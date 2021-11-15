<script>
	import Entry from '$lib/Entry.svelte';
  import { onMount } from 'svelte';

	export let entries = [];

  let limit = 0;
  let limitedEntries = [];
  let i = false;
	let noEntries = false;

  $: reload(entries);

  function reload() {
    if (!i) return;
    //console.log('reload!')
    const h = document.documentElement.scrollHeight;
    const n = parseInt(h / 120); // 120 = min. entry height
    limit = n;
    limitedEntries = entries.slice(0, limit);
		if (limitedEntries.length < 1) noEntries = true;
  }

  function init() {
    i = true;
    reload();

    window.addEventListener('scroll',() => {
      const { scrollHeight, scrollTop, clientHeight } = document.documentElement;
      if (scrollTop + clientHeight > scrollHeight - 25) {
        //console.log('load more!')
        limit += 5;
        limitedEntries = entries.slice(0, limit);
    	}
    });
  }

	onMount(async () => {
		init();
	});
</script>

<div class="entries">
	{#if noEntries}
		<small class="info">No entries found...</small>
	{:else}
		{#each limitedEntries as entry}
			<Entry {entry} />
		{:else}
			<small class="info">
				loading...
			</small>
		{/each}
	{/if}
</div>

<style>
</style>
