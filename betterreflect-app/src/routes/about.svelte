<script>
	import SideNav from '$lib/SideNav.svelte';
	import Main from '$lib/Main.svelte';
</script>

<div class="sidenav">
	<SideNav entries={[]} on:change={e => navChange(e.detail)} />
</div>
<Main entries={[]} showSideNav={false} type="about" />

<style>
	.sidenav {
		display: none;
	}
	@media all and (max-width: 600px) { /* 1000px = side width + max. main width */
		.sidenav {
			display: initial;
		}
	}
</style>
