<script context="module">
  export async function load({ page, fetch, session, context }) {

  	const url = '/entry/sessions.json';

  	const res = await fetch(url);
  	if (res.ok) {
  		return {
  			props: {
  				sessions: sessions
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
  import Main from '$lib/Main.svelte';

  export let entry = {};
  export let edit = false;

  let showSideNav = true;
</script>

<SideNav entries={[]} hidden={showSideNav} backbutton={true} ref={entry.type} />
{#if edit}
  <Main {entry} type="edit" />
{:else}
  <Main {entry} type="entry" />
{/if}

<style>
</style>
