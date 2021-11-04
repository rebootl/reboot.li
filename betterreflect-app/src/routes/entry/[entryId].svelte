<script context="module">
  export async function load({ page, fetch, session, context }) {

    const entryId = page.params.entryId;
  	const url = `/entries/${entryId}.json`;
    const edit = page.query.has('edit');

		console.log('load task entry:')
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

</script>

<div class="wrapper">
	<aside>
		<!--<SideNav />-->
		test
	</aside>

	<main>
		<h1>tasks</h1>
	</main>
</div>

<style>
	.wrapper {
		display: flex;
	}
	aside {
		width: var(--side-width);
	}
</style>
