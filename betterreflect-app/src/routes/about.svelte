<script>
	import SideNav from '$lib/SideNav.svelte';
</script>

<div class="sidenav">
	<SideNav entries={[]} on:change={e => navChange(e.detail)} />
</div>
<main>
	<h1>About</h1>

	This is a personal website.
</main>

<style>
	main {
		box-sizing: border-box;
		margin-left: auto;
		margin-right: auto;
		max-width: var(--max-main-width);
		padding: 0 20px 30px 20px;
	}
	.sidenav {
		display: none;
	}
	@media all and (max-width: 600px) { /* 1000px = side width + max. main width */
		main {
			margin-left: initial;
			margin-right: initial;
		}
		.sidenav {
			display: initial;
		}
	}
</style>
