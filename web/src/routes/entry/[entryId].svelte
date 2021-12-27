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

  export let entry = {};
  export let edit = false;

  let showSideNav = true;
</script>

<SideNav entries={[]} hidden={showSideNav} backbutton={true} ref={entry.type} />
<main class="margin-left">
  {#if edit}
    <EditEntry {entry} ref={entry.type} />
  {:else}
    <Entry {entry} />
  {/if}
</main>
