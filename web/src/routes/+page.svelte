<script>
	import SideNav from '$lib/SideNav.svelte';
	import NewEntry from '$lib/NewEntry.svelte';
	import EntriesFilter from '$lib/EntriesFilter.svelte';
	/*import { session } from '$app/stores';*/
	/*import { page } from '$app/stores';*/

	export let data;
	let entries = data.entries;
	//console.log(data)

	$: init(entries);

	function init() {
		entries = entries.slice(0, 3);
	}

	function created(e) {
	  entries.push(e);
	  entries = entries;
	}
</script>

<!--<SideNav entries={[]} showOnWide={false} />-->
<main>
	<h1>Welcome!</h1>

	Browse my entries:

	<ul>
	  <li><a href="/tasks">Tasks</a>
	  <li><a href="/notes">Notes</a>
	  <li><a href="/links">Links</a>
	  <li><a href="/images">Images</a>
	</ul>

	or visit the <a href="/about">About</a> page.

	<h2>News</h2>
	{#if data.loggedIn}
		<NewEntry type={'news'} on:created={e => created(e.detail)} />
	{/if}
	<div class="newsbox">
		<EntriesFilter {entries} />
		<div class="newslink"><a href="/news"><small>View all</small></a></div>
	</div>
</main>

<style>
	.newsbox {
		padding: 0 20px 20px 20px;
    border: 1px solid var(--main-line-color);
	}
	.newslink {
		display: flex;
		justify-content: center;
    padding-top: 15px;
	}
</style>
