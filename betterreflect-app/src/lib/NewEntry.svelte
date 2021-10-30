<script>
  import EditTypes from './EditTypes.svelte';
  import EditTopics from './EditTopics.svelte';
  import EditTags from './EditTags.svelte';
  import LoadImages from './LoadImages.svelte';
  import { compressImage, encodeData, uploadMultiImagesGenerator }
    from '$lib/images';
  import { session } from '$app/stores';
  import { createEventDispatcher } from 'svelte';
  import { goto } from '$app/navigation';

  const dispatch = createEventDispatcher();

  export let topics = [];
  export let tagsByTopics = {};
  export let edit = false;
  export let entry = {};

  let showAddElements = true;

  let text = '';
  let type = 'task';
  let newTopics = [];
  let newTags = [];
  let _private = false;
  let pinned = false;
  let linkComment = '';
  let linkTitle = '';
  let images = [];
  let newImages = [];
  let resetLoadImages = [];

  let loadTopics = [];
  let loadTags = [];

  $: textInput(text)
  $: loadEntry(entry);

  function loadEntry() {
    if (!edit) return;
    if (!entry) return;
    type = entry.type;
    if (type === 'link') {
      linkComment = entry.commment;
      linkTitle = entry.title;
    }
    text = entry.text;
    loadTopics = entry.topics;
    newTopics = entry.topics;
    loadTags = entry.tags;
    newTags = entry.tags;
    _private = entry.private;
    pinned = entry.pinned;
    showAddElements = true;
  }

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

  async function uploadNewImages() {
    let uploadResult = {};
    console.log(newImages)
    for await (const r of uploadMultiImagesGenerator(newImages)) {
      // update progress
      uploadResult = r;
      //uploadProgress = r.progress;
    }
    // handle the upload result
    if (!uploadResult.result.success) return false;
    newImages.forEach(i => {
      const r = uploadResult.result.files.find(e => e.originalname === i.filename);
      i.filepath = r.path;
      delete i.file;
      return i;
    });
    return true;
  }

  async function sendRequest(type, data) {
    let r;
    try {
      const res = await fetch(`/entries/${$session.user}.json`, {
        method: 'PUT',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
      });
      if (res.ok) {
        r = await res.json();
        return r;
      } else {
        const { message } = await res.json();
        new Error(message);
      }
    } catch(error) {
      console.error(error);
      return {
        success: false,
        result: error
      }
    }
  }

  async function create() {
    if (newTopics.length < 1) {
      console.log('at least one topic must be selected')
      return;
    }
    console.log('create!')

    // upload new images
    if (newImages.length > 0) {
      const res = await uploadNewImages();
      if (!res) {
        console.log('error at uploading images')
        return;
      }
    }

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
      entry.images = newImages;
    }

    //const r = await sendRequest('POST', entry);
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
    if (!r.success) {
      console.log('error creating entry');
      return;
    }

    console.log('success!')
    dispatch('created', entry);
    reset();
  }

  async function update() {
    if (newTopics.length < 1) {
      console.log('at least one topic must be selected')
      return;
    }
    console.log('update!')

    entry.mdate = new Date();
    entry.type = type;
    entry.private = _private;
    entry.pinned = pinned;
    entry.topics = newTopics;
    entry.tags = newTags;
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
        method: 'PUT',
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
    if (!r.success) {
      console.log('error updating entry');
      return;
    }

    console.log('success!')
    dispatch('updated', r);
    goto(`/entries/${$session.user}`);
  }

  async function _delete() {
    if (!confirm("Do u really want to delete this entry?"));
      return;

    let r;
    try {
      const res = await fetch(`/entries/${$session.user}.json`, {
        method: 'DELETE',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(entry)
      });
      if (res.ok) {
        console.log(res)
        r = await res.json();
      } else {
        const { message } = await res.json();
        new Error(message);
      }
    } catch(error) {
      console.error(error);
    }
    if (!r.success) {
      console.log('error deleting entry');
      return;
    }

    console.log('success!')
    dispatch('deleted', r);
    goto(`/entries/${$session.user}`);
  }

  function loadImages(images) {
    if (images.length > 0) {
      newImages = images;
      type = 'image';
      showAddElements = true;
    } else {
      newImages = images;
      type = 'task';
      showAddElements = false;
    }
  }

  function reset() {
    text = '';
    type = 'task';
    newTopics = [];
    newTags = [];
    loadTopics = [];
    loadTags = [];
    _private = false;
    pinned = false;
    linkComment = '';
    linkTitle = '';
    images = [];
    newImages = [];
    showAddElements = false;
    resetLoadImages = [];
  }

</script>

<div class="newentry-box">

  <div>
    {#if newImages.length <= 0}
      <textarea class="newentry-text"
                placeholder="New Entry..."
                bind:value={text}></textarea>
    {/if}
    <LoadImages on:change={(e) => loadImages(e.detail)}
                reset={resetLoadImages} />
  </div>
  {#if showAddElements}
    {#if type === 'image'}
      Type: Image
    {:else}
      <EditTypes selectedType={edit ? entry.type : 'task'}
                 on:change={(e) => setType(e.detail)} />
    {/if}
    {#if type === 'link'}
      <input id="linktitle" name="linktitle" placeholder="Link title..."
             bind:value={linkTitle}>
      <input id="linkcomment" name="linkcomment" placeholder="Link comment..."
             bind:value={linkComment}>
    {/if}
    <EditTopics items={topics} selectedItems={loadTopics}
                on:change={(e) => setNewTopics(e.detail)} />
    <EditTags {tagsByTopics} {newTopics} {loadTopics}
              selectedItems={loadTags}
              on:change={(e) => setNewTags(e.detail)} />
    <div>
      <input type="checkbox" id="private-checkbox" name="private"
             bind:checked={_private}>
      <label for="private-checkbox">Private</label>
      <input type="checkbox" id="pinned-checkbox" name="pinned"
             bind:checked={pinned}>
      <label for="pinned-checkbox">Pinned</label>
    </div>
    <div>
      {#if edit}
        <div class="editbuttons">
          <div>
            <button on:click={() => update()}>Update</button>
            <a href={'/entries/' + $session.user} class="cancelbutton">
              <small>Cancel</small></a>
          </div>
          <button on:click={() => _delete()} class="deletebutton">
            Delete
          </button>
        </div>
      {:else}
        <button on:click={() => create()}>Create</button>
      {/if}
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
  .editbuttons {
    display: flex;
    justify-content: space-between;
  }
  .cancelbutton {
    margin-left: 10px;
  }
  .deletebutton {
    background-color: var(--error-color);
    color: var(--on-error-color);
    border: 2px solid var(--error-color);
    border-radius: 5px;
  }
</style>
