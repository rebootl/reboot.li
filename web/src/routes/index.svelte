<script>
	import SideNav from '$lib/SideNav.svelte';
</script>

<SideNav entries={[]} showOnWide={false} />
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
</main>
