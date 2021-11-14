
<script>
  import NewEntry from '$lib/NewEntry.svelte';
  import Entries from '$lib/Entries.svelte';
  import Home from '$lib/Home.svelte';
	import About from '$lib/About.svelte';
  import Entry from '$lib/Entry.svelte';
  import EditEntry from '$lib/EditEntry.svelte';
  import { session } from '$app/stores';
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let type = '';
  export let showSideNavOnWide = true;
	export let entries = [];
  export let entry = {};

</script>

<main class:margin-left={showSideNavOnWide}>
  {#if type === 'home'}
    <Home />
  {:else if type === 'about'}
    <About />
  {:else if type === 'entry'}
    <Entry {entry} />
  {:else if type === 'edit'}
    <EditEntry {entry} ref={entry.type} />
  {:else}
    {#if $session.loggedIn}
      <NewEntry type={type} on:created={e => dispatch('created', e.detail)}/>
    {/if}
    <Entries entries={entries} />
  {/if}
</main>

<style>
  main {
    box-sizing: border-box;
    max-width: var(--max-main-width);
    padding: 0 20px 30px 20px;
    margin-left: auto;
    margin-right: auto;
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
  @media all and (max-width: 600px) { /* 1000px = side width + max. main width */
    .margin-left {
      margin-left: initial;
    }
  }
  .info {
    padding-top: 10px;
  }
</style>
