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

  export let entry = {};
  export let edit = false;

  let showSideNav = true;

</script>

<SideNav entries={[]} hidden={showSideNav} backbutton={true} ref={entry.type} />
<main class:margin-left={showSideNav} >
  <Entry {entry} />
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
