<script context="module">

  export async function load({ page, fetch, session, context }) {

    const user = page.params.user;
  	const url = `/entries/${user}.json`;
    const entryId = page.query.has('entryId') ? page.query.get('entryId') : '';

  	const res = await fetch(url);

  	if (res.ok) {
  		return {
  			props: {
          user: user,
  				entries: await res.json(),
          entryId: entryId,
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
  import Entry from '$lib/Entry.svelte';
  import NewEntry from '$lib/NewEntry.svelte';
  import Topics from '$lib/Topics.svelte';
  import Tags from '$lib/Tags.svelte';
  import Types from '$lib/Types.svelte';
  import IconButton from '$lib/IconButton.svelte';
  import BackButton from '$lib/BackButton.svelte';
  import { session } from '$app/stores';

  export let user = '';
  export let entries = [];
  export let entryId = '';

  let loggedIn = $session.loggedIn;

  let filteredEntries = entries;

  let topics = [];
  let tags = [];
  let tagsByTopics = {};
  let selectedTopics = new Set();
  let selectedTags = new Set();
  let selectedType = 'any';
  let showPrivate = true;
  let showPinned = true;
  let entryNotFound = false;
  let singleEntry = {};

  $: setTopics(entries);
  $: filterEntries(selectedTopics, selectedTags, selectedType, showPrivate,
      showPinned);

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
    if (entryId === '') {
      singleEntry = {};
      entryNotFound = false;
      filteredEntries = getFilteredEntries();
    } else {
      singleEntry = entries.find(e => e.id === entryId);
      if (!singleEntry) entryNotFound = true;
      else entryNotFound = false;
      filteredEntries = [];
    }
  }

  function getFilteredEntries() {
    let f = entries;

    if (selectedType !== 'any') {
      f = entries.filter(e => e.type === selectedType);
    } /*else {
      f = entries;
    }*/

    //console.log(selectedTopics)
    if (selectedTopics.size > 0) {
      f = f.filter(e => {
        for (const t of e.topics) {
          if (selectedTopics.has(t)) return e;
        }
      });
    }

    //console.log(selectedTags)
    if (selectedTags.size > 0) {
      f = f.filter(e => {
        for (const t of e.tags) {
          if (selectedTags.has(t)) return e;
        }
      });
    }

    // sort / pinned
    let r;
    if (showPinned) {
      const p = f.filter(e => e.pinned).sort((a, b) => a.date - b.date);
      const q = f.filter(e => !e.pinned).sort((a, b) => a.date - b.date);
      r = [ ...p, ...q ];
    } else {
      r = f.sort((a, b) => a.date - b.date);
    }

    // filter private
    if (showPrivate) {
      return r;
    } else {
      return r.filter(e => !e.private);
    }
  }

</script>

<nav class="sidenav">
  {#if entryId === ''}
    <div class="back-button">
      <BackButton href='/' icon='home'>Home</BackButton>
    </div>
    <h2 class="username">{user}</h2>
    <Topics {topics} on:change={(e) => selectTopic(e.detail)} />
    <Tags {tags} on:change={(e) => selectedTags = e.detail} />
  {:else}
    <div class="back-button">
      <BackButton href={'/entries/' + user} icon="person">{user}</BackButton>
    </div>
    <!--<h2 class="username">Entry</h2>-->
  {/if}
</nav>

<main class="main">
  {#if entryId === ''}
    {#if loggedIn}
      <NewEntry {topics} {tagsByTopics} />
    {/if}
    <div class="entry-filters">
      <div class="shownav">
        <div>
          <input type="checkbox" id="show-pinned" name="show-pinned" checked
                 on:click={() => showPinned = !showPinned}>
          <label for="show-pinned">Show pinned on top</label>
        </div>
        {#if loggedIn}
          <div>
            <input type="checkbox" id="show-private" name="show-private" checked
                   on:click={() => showPrivate = !showPrivate}>
            <label for="show-private">Show private entries</label>
          </div>
        {/if}
      </div>
      <div class="typenav">
        <Types on:change={(e) => selectedType = e.detail} />
      </div>
    </div>
    <Entries entries={filteredEntries} />
  {:else}
    {#if entryNotFound}
      Oops, entry not found :(
    {:else}
      <Entry entry={singleEntry} />
    {/if}
  {/if}
</main>

<style>
  .sidenav {
    position: fixed;
    width: 220px;
    height: calc(100% - var(--header-height));
    overflow-y: scroll;
  }
  .back-button {
    margin-top: 5px;
  }
  .username {
    padding-left: 10px;
  }
  .main {
    margin-left: 220px;
    padding: 0 20px 0 20px;
  }
  .entry-filters {
    display: flex;
    flex-flow: column;
    gap: 10px;
    margin: 20px 0 20px 0;
  }
  .shownav {
    display: flex;
    gap: 10px;
    flex-wrap: wrap;
  }
</style>
