<script context="module">
  export async function load({ page, fetch, session, context }) {

  	const url = '/entries/task.json';
    const edit = page.query.has('edit');

		console.log('load tasks')
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

</script>

<div class="wrapper">
	<SideNav {entries} />

	<main>
		{#each entries as entry}
			<Entry {entry} />
		{/each}
	</main>
</div>

<style>
	.wrapper {
		display: flex;
	}
	main {
		width: 100%;
		height: calc(100vh - var(--header-height));
    overflow-y: scroll;
    padding: 0 20px 0 20px;
  }
</style>
