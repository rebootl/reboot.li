<script>
  import moment from 'moment';

  import EditTopics from './EditTopics.svelte';
  import EditTags from './EditTags.svelte';
  import LoadImages from './LoadImages.svelte';
  import { sendRequest, sendTokenRequest, getToken } from '$lib/request';
  import { currentTopics, currentTags, currentTagsByTopics } from '$lib/store';
  import { refs } from '$lib/refs';
  import { MEDIASERVER } from '../../config.js';

  import { compressImage, encodeData, uploadMultiImagesGenerator }
    from '$lib/images';

  import { session } from '$app/stores';
  import { createEventDispatcher } from 'svelte';
  import { goto } from '$app/navigation';

  const dispatch = createEventDispatcher();

  export let entry = {};
  export let ref = '';

  let type = '';
  let text = '';
  let newTopics = [];
  let newTags = [];
  let _private = false;
  let pinned = false;
  let linkComment = '';
  let linkTitle = '';
  let imagesDate = '';
  let images = [];
  let newImages = [];
  let resetLoadImages = [];

  let loadTopics = [];
  let loadTags = [];

  $: loadEntry(entry);

  function loadEntry() {
    if (!entry) return;
    type = entry.type;
    if (type === 'link') {
      linkComment = entry.commment;
      linkTitle = entry.title;
    }
    if (type === 'image') {
      const m = moment(entry.imagesDate);
      imagesDate = m.isValid() ? m.format('YYYY-MM-DD') : '';
    }
    text = entry.text;
    loadTopics = entry.topics;
    newTopics = entry.topics;
    loadTags = entry.tags;
    newTags = entry.tags;
    _private = entry.private;
    pinned = entry.pinned;
  }

  function setNewTopics(v) {
    newTopics = v;
  }

  function setNewTags(v) {
    newTags = v;
  }

  async function uploadNewImages() {
    // get mediaserver token
    const token = await getToken();
    if (!token) {
      console.log('error getting mediaserver token');
      return;
    }

    // (this is here for eventual progress indicator, not used yet)
    // (and also handling upload result)
    let uploadResult = {};
    for await (const r of uploadMultiImagesGenerator(newImages, token)) {
      // update progress
      uploadResult = r;
      //uploadProgress = r.progress;
    }
    // handle the upload result
    if (!uploadResult.result?.success) return false;
    newImages.forEach(i => {
      const r = uploadResult.result.files.find(e => e.originalname === i.filename);
      i.filepath = r.path;
      i.url = r.url;
      delete i.file;
      delete i.maxSize;
      return i;
    });
    return true;
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
    if ([ 'task', 'note', 'link' ].includes(entry.type)) {
      entry.text = text;
    } else if (type === 'image') {
      // -> upload new images
      if (newImages.length > 0) {
        const res = await uploadNewImages();
        if (!res) {
          console.log('error at uploading images');
          return;
        }
      }
      const m = moment(imagesDate);
      entry.imagesDate = m.isValid() ? m.toDate() : '';
      entry.images = [ ...entry.images, ...newImages ];
    }
    if (entry.type === 'link') {
      entry.comment = linkComment;
      entry.title = linkTitle;
    }
    const r = await sendRequest('PUT', `/entry/${entry.id}.json`, entry);
    if (!r.success) {
      console.log('error updating entry');
      return;
    }

    console.log('success!')
    dispatch('updated', r);
    goto(refs[ref].href);
  }

  async function _delete() {
    if (!confirm("Do u really want to delete this entry?"));
      return;

    if (type === 'image') {
      // get mediaserver token
      const token = await getToken();
      if (!token) {
        console.log('error getting mediaserver token');
        return;
      }

      for (const i of entry.images) {
        const r = await sendTokenRequest('POST',
          new URL('/api/deleteImage', MEDIASERVER),
          { filepath: i.filepath },
          token
        );
        if (!r.success) {
          console.log('error deleting image');
          return;
        }
      }
    }

    const r = await sendRequest('DELETE', `/entry/${entry.id}.json`, entry);
    if (!r.success) {
      console.log('error deleting entry');
      return;
    }

    console.log('success!')
    dispatch('deleted', r);
    goto(refs[ref].href);
  }

  function loadNewImages(v) {
    newImages = v;
  }

  async function deleteImage(image) {
    if (!confirm("Do u really want to delete this image?"));
      return;
    console.log('deleting image');

    // get mediaserver token
    const token = await getToken();
    if (!token) {
      console.log('error getting mediaserver token');
      return;
    }

    const r = await sendTokenRequest('POST',
      new URL('/api/deleteImage', MEDIASERVER),
      { filepath: image.filepath },
      token
    );
    if (!r.success) {
      console.log('error deleting image');
      return;
    }
    entry.images = entry.images.filter(i => i.filepath !== image.filepath);

    // update image array on server
    const s = await sendRequest('PUT', `/entry/${entry.id}.json`, entry);
    if (!s.success) {
      console.log('error updating entry');
      return;
    }

    console.log('sucess!');
  }

  /*function reset() {
    text = '';
    type = '';
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
    resetLoadImages = [];
  }*/

</script>

<div class="edit-box">
    {#if type === 'image'}
      <LoadImages on:change={(e) => loadNewImages(e.detail)}
                  reset={resetLoadImages} />
    {:else if type === 'link'}
      <input class="link-text"
             placeholder="Edit Entry..."
             bind:value={text}>
    {:else}
      <textarea class="edit-text"
                placeholder="Edit Entry..."
                bind:value={text}></textarea>
    {/if}
  {#if type === 'image'}
    <div class="editimages">
      {#if entry.images}
        {#each entry.images as image}
          <img class="editimage" src={image.url} alt="edit preview" />
          <textarea class="imagecomment" bind:value={image.comment}
                    placeholder="Comment..."></textarea>
          <button on:click={e => deleteImage(image)}
                  class="deletebutton">Delete</button>
        {/each}
      {/if}
      <label for="imagesdate">Images date (manual):</label>
      <input id="imagesdate" name="imagesdate" type="date"
             bind:value={imagesDate}>
    </div>
  {/if}
  {#if type === 'link'}
    <input id="linktitle" name="linktitle" placeholder="Link title..."
           bind:value={linkTitle}>
    <input id="linkcomment" name="linkcomment" placeholder="Link comment..."
           bind:value={linkComment}>
  {/if}
  <EditTopics items={$currentTopics} selectedItems={loadTopics}
              on:change={(e) => setNewTopics(e.detail)} />
  <EditTags tagsByTopics={$currentTagsByTopics} {newTopics} {loadTopics}
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
  <div class="editbuttons">
    <div>
      <button on:click={() => update()}>Update</button>
      <a href={refs[ref].href} class="cancelbutton">
        <small>Cancel</small></a>
    </div>
    <button on:click={() => _delete()} class="deletebutton">
      Delete
    </button>
  </div>
</div>

<style>
  .edit-box {
    display: flex;
    flex-flow: column;
    gap: 20px;
    padding: 20px 0 20px 0;
    border-bottom: 1px solid var(--main-line-color);
  }
  .edit-text {
    width: 500px;
    height: 300px;
    padding: 10px;
  }
  .editimages {
    display: flex;
    align-items: flex-start;
    flex-flow: column;
    gap: 20px;
    padding-bottom: 20px;
    border-bottom: 1px solid var(--main-line-color);
  }
  .editimage {
    max-width: 240px;
    max-height: 240px;
  }
  .editbuttons {
    display: flex;
    align-self: stretch;
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
