<script context="module">
	/** @type {import('@sveltejs/kit').ErrorLoad} */
	export function load({ error, status }) {
		return {
			props: {
				status: status,
        message: error.message
			}
		};
	}
</script>

<script>
  import SideNav from '$lib/SideNav.svelte';

  export let status = '';
	export let message = '';
</script>

<SideNav entries={[]} showOnWide={false} />
<main>
  <h1>Oops error: {status}</h1>
  <small>{message}</small>
</main>
