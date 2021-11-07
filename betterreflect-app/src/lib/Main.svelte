
<script>
  import NewEntry from '$lib/NewEntry.svelte';
	import Entries from '$lib/Entries.svelte';
  import { getFilteredEntries } from '$lib/filterSort';
  import { session } from '$app/stores';

  export let type = '';
  export let showSideNav = true;
	export let entries = [];

  let filteredEntries = [];

  $: filterEntries([[], []], entries);

  function filterEntries(v) {
    filteredEntries = getFilteredEntries(entries, v);
  }
</script>

<main class:margin-left={showSideNav}>
  {#if $session.loggedIn}
    <NewEntry type={type} />
  {/if}
  <Entries entries={filteredEntries} {showSideNav} />
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
