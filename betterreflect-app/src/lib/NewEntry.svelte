<script>
  import EditTypes from './EditTypes.svelte';
  import EditTopics from './EditTopics.svelte';
  import EditTags from './EditTags.svelte';
  import { session } from '$app/stores';
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let topics = [];
  export let tagsByTopics = {};

  let showAddElements = true;

  let text = '';
  let type = 'task';
  let newTopics = [];
  let newTags = [];
  let _private = false;
  let pinned = false;
  let linkComment = '';
  let linkTitle = '';

  $: textInput(text)

  function textInput() {
    if (text === '') showAddElements = false;
    else showAddElements = true;
  }

  function setType(v) {
    type = v;
  }

  function setNewTopics(v) {
    newTopics = v;
  }

  function setNewTags(v) {
    newTags = v;
  }

  async function create() {
    if (newTopics.length < 1) return;
    console.log('create!')

    const entry = {
      id: type + '-' + Date.now().toString(36) +
                 Math.random().toString(36).substr(2, 5),
      date: new Date(),
      user: $session.user,
      type: type,
      topics: newTopics,
      tags: newTags,
      private: _private,
      pinned: pinned,
    }

    if ([ 'task', 'article', 'link' ].includes(entry.type)) {
      entry.text = text;
    } else if (entry.type === 'link') {
      entry.comment = linkComment;
      entry.title = linkTitle;
    } else if (type === 'image') {
      //d.images = [ ...images, ...newImages ];
    }

    let r;
    try {
      const res = await fetch(`/entries/${$session.user}.json`, {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(entry)
      });
      if (res.ok) {
        r = await res.json();
      } else {
        const { message } = await res.json();
        new Error(message);
      }
    } catch(error) {
      console.error(error);
    }
    if (!r) return;

    console.log('success!')
    dispatch('created', r);
    reset();
  }

  function reset() {
    text = '';
    type = 'task';
    newTopics = [];
    newTags = [];
    _private = false;
    pinned = false;
    linkComment = '';
    linkTitle = '';
  }

</script>

<div class="newentry-box">

  <div>
    <textarea class="newentry-text"
              placeholder="New Entry..."
              bind:value={text}></textarea>
  </div>
  {#if showAddElements}
    <EditTypes on:change={(e) => setType(e.detail)} />
    <EditTopics items={topics}
                on:change={(e) => setNewTopics(e.detail)} />
    <EditTags {tagsByTopics} {newTopics} on:change={(e) => setNewTags(e.detail)} />
    <div>
      <input type="checkbox" id="private-checkbox" name="private"
             bind:checked={_private}>
      <label for="private-checkbox">Private</label>
      <input type="checkbox" id="pinned-checkbox" name="pinned"
             bind:checked={pinned}>
      <label for="pinned-checkbox">Pinned</label>
    </div>
    <div>
      <button on:click={() => create()}>Create</button>
    </div>
  {/if}

</div>

<style>
  .newentry-box {
    display: flex;
    flex-flow: column;
    gap: 20px;
    padding: 20px 0 20px 0;
    border-bottom: 1px solid var(--main-line-color);
  }
  .newentry-text {
    width: 170px;
    height: 20px;
    padding: 10px;
  }
</style>
