<script context="module">
  export async function load({ page, fetch, session, context }) {

  	const url = '/entries/link.json';
    const edit = page.query.has('edit');

		console.log('load links')
  	const res = await fetch(url);
  	if (res.ok) {
  		return {
  			props: {
  				entries: await res.json(),
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

	export let entries = [];
  export let showSideNav = true;

  let filteredEntries = entries;

  function filterEntries(v) {
    const [ selectedTopics, selectedTags ] = v;
    console.log(selectedTopics)
    let f = entries;
    if (selectedTopics.length > 0) {
      f = f.filter(e => {
        for (const t of e.topics) {
          if (selectedTopics.includes(t)) return e;
        }
      });
      console.log(f)
    }
    if (selectedTags.length > 0) {
      f = f.filter(e => {
        for (const t of e.tags) {
          if (selectedTags.includes(t)) return e;
        }
      });
    }
    filteredEntries = f;
  }

</script>

<div class="wrapper">
	<SideNav {entries} hidden={showSideNav}
           on:change={e => filterEntries(e.detail)} />

	<main class:margin-left={showSideNav} >
		{#each filteredEntries as entry}
			<Entry {entry} />
		{/each}
	</main>
</div>

<style>
	.wrapper {
	}
	main {
    display: flex;
    flex-flow: column;
		min-height: calc(100vh - var(--header-height));
    padding: 0 20px 0 20px;
    overflow: hidden;
    max-width: var(--max-main-width);
  }
  .margin-left {
    margin-left: var(--side-width);
  }
</style>
