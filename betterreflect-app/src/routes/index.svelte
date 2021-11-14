<script>
	import SideNav from '$lib/SideNav.svelte';
	import Main from '$lib/Main.svelte';
</script>

<SideNav entries={[]} on:change={e => navChange(e.detail)} showOnWide={false} />
<Main entries={[]} showSideNavOnWide={false} type="home" />

<style>
</style>
