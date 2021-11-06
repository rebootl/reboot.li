<script>
	import Entry from '$lib/Entry.svelte';
  import { onMount } from 'svelte';

	export let entries = [];
  export let showSideNav = true;

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

<main class:margin-left={showSideNav} >
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
</main>

<style>
	main {
    display: flex;
    flex-flow: column;
		/*min-height: calc(100vh - var(--header-height));*/
    padding: 0 20px 30px 20px;
    overflow: hidden;
    max-width: var(--max-main-width);
  }
  .margin-left {
    margin-left: var(--side-width);
  }
	@media all and (min-width: 1000px) { /* 1000px = side width + max. main width */
		.margin-left {
			margin-left: auto;
	    margin-right: auto;
	  }
	}
	.info {
		padding-top: 10px;
	}
</style>
