<script context="module">

  export async function load({ page, fetch, session, context }) {
    console.log(page)
    const user = page.params.user;
  	const url = `/entries/${user}.json`;

  	const res = await fetch(url);

  	if (res.ok) {
  		return {
  			props: {
          user: user,
  				entries: await res.json()
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
  import Entries from '$lib/Entries.svelte';

  export let user = '';
  export let entries = [];

  let topics = [];

  $: setTopics(entries);

  // derived from userEntries
  function setTopics() {
    const t = [];
    for (const e of entries) {
      for (const topic of e.topics) {
        if (!t.includes(topic)) {
          t.push(topic);
        }
      }
    }
    topics = t.sort();
    console.log(topics)
  }

</script>

<h2>User: {user}</h2>

<Entries {entries} />

<style>

</style>
