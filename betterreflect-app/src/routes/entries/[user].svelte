<script context="module">

  export async function load({ page, fetch, session, context }) {
    //console.log(page)
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
  import Topics from '$lib/Topics.svelte';
  import Tags from '$lib/Tags.svelte';

  export let user = '';
  export let entries = [];

  let filteredEntries = entries;

  let topics = [];
  let tags = [];
  let tagsByTopics = {};
  let selectedTopics = new Set();
  let selectedTags = new Set();

  $: setTopics(entries);
  $: filterEntries(selectedTopics, selectedTags);

  // create topics and tagsByTopics from entries
  function setTopics() {
    const t = [];

    entries.forEach(entry => {
      entry.topics.forEach(topic => {
        if (!t.includes(topic)) {
          t.push(topic);
        }

        if (!tagsByTopics.hasOwnProperty(topic))
          // must create copy here!!
          tagsByTopics[topic] = [ ...entry.tags ];
        else
          entry.tags.forEach(tag => {
            if (!tagsByTopics[topic].includes(tag))
              tagsByTopics[topic].push(tag);
          });
      });
    });

    topics = t.sort();
    setTags();
  }

  function selectTopic(v) {
    selectedTopics = v;
    setTags();
  }

  function setTags() {
    selectedTags.clear();
    selectedTags = selectedTags;

    const r = [];
    if (selectedTopics.size > 0) {
      selectedTopics.forEach(topic => {
        tagsByTopics[topic].forEach(tag => {
          if (!r.includes(tag)) r.push(tag);
        });
      });
    } else {
      entries.forEach(e => {
        e.tags.forEach(tag => {
          if (!r.includes(tag)) r.push(tag);
        });
      });
    }

    tags = r.sort();
  }

  function filterEntries() {
    let f = entries;
    let t = [];
    //console.log(selectedTopics)
    if (selectedTopics.size > 0) {
      f = entries.filter((e) => {
        for (const t of e.topics) {
          if (selectedTopics.has(t)) return e;
        }
      });
    }

    //console.log(selectedTags)
    if (selectedTags.size > 0) {
      f = f.filter((e) => {
        for (const t of e.tags) {
          if (selectedTags.has(t)) return e;
        }
      });
    }

    filteredEntries = f;
  }

</script>

<h2>User: {user}</h2>

<Topics {topics} on:change={(e) => selectTopic(e.detail)} />
<Tags {tags} on:change={(e) => selectedTags = e.detail} />

<Entries entries={filteredEntries} />

<style>

</style>
