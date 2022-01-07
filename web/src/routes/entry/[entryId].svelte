<script context="module">
  export async function load({ page, fetch, session, context }) {

    const entryId = page.params.entryId;
  	const url = `/entry/${entryId}.json`;
    const edit = page.query.has('edit');

		console.log('load entry:')
    console.log(entryId)

  	const res = await fetch(url);
  	if (res.ok) {
  		return {
  			props: {
  				entry: await res.json(),
          edit: edit,
  			}
  		};
  	}

  	return {
  		status: res.status,
  		error: new Error(`Could not load ${url}`)
  	};
  }
</script>

<script>
  import SideNav from '$lib/SideNav.svelte';
  import Entry from '$lib/Entry.svelte';
  import EditEntry from '$lib/EditEntry.svelte';
  import { session } from '$app/stores';

  export let entry = {};
  export let edit = false;

  let showSideNav = true;
  let versions = []
  let initDone = false;

  $: init(entry);

  function init() {
    if (initDone) return;
    for (let v = 1; v <= entry.version; v++) {
      versions.push(v);
    }
    versions = versions;
    initDone = true;
  }

  async function getVersion(v) {
    const url = `/getVersion/${entry.id}/${v}.json`;
    try {
      const r = await fetch(url);
      if (r.ok) {
        entry = await r.json();
      } else {
        throw(new Error(`${r.status} ${r.statusText}`));
      }
    } catch(err) {
      console.error(err);
    }
  }
</script>

<SideNav entries={[]} hidden={showSideNav} backbutton={true} ref={entry.type} />
<main class="margin-left">
  {#if $session.loggedIn && !edit}
    <div class="versions">
      {#each versions as v}
        {#if v === entry.version}
          <div class="version current">
            {v}
          </div>
        {:else}
          <div class="version" on:click={() => getVersion(v)}>
            {v}
          </div>
        {/if}
      {/each}
    </div>
  {/if}
  {#if edit}
    <EditEntry {entry} ref={entry.type} />
  {:else}
    <Entry {entry} />
  {/if}
</main>

<style>
  .versions {
    display: flex;
    gap: 10px;
  }
  .version {
    padding: 5px 10px 5px 10px;
    border: 1px solid var(--main-line-color);
  }
  .version.current {
    background-color: var(--main-highlight-background-color);
  }
</style>
