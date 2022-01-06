<script>
  import moment from 'moment';

  import EditTopics from './EditTopics.svelte';
  import EditTags from './EditTags.svelte';
  import LoadImages from './LoadImages.svelte';
  import { sendRequest, getToken } from '$lib/request';
  import { debounce } from '$lib/helper';
  import { currentTopics, currentTags, currentTagsByTopics } from '$lib/store';

  import { compressImage, encodeData, uploadMultiImagesGenerator }
    from '$lib/images';

  import { session } from '$app/stores';
  import { createEventDispatcher } from 'svelte';
  import { goto } from '$app/navigation';

  const dispatch = createEventDispatcher();

  export let type = 'task';

  let showAddElements = false;

  let text = '';
  let newTopics = [];
  let newTags = [];
  let _private = false;
  let pinned = false;
  let linkComment = '';
  let linkTitle = '';
  let linkTitleDisabled = true;
  let linkErr = false;
  let linkErrCode = '';
  let images = [];
  let newImages = [];
  let resetLoadImages = [];

  $: textInput(text)

  function textInput() {
    if (text === '') showAddElements = false;
    else showAddElements = true;

    if (type === 'link' && text !== '') {
      linkTitleDisabled = true;
      linkTitle = 'getting title...';
      debounce(() => getTitle(text), 500);
    }
  }

  function setNewTopics(v) {
    newTopics = v;
  }

  function setNewTags(v) {
    newTags = v;
  }

  async function getTitle(text) {
    //console.log('getTitle')
    const r = await sendRequest('POST', '/getTitle.json', {
      url: text
    });
    if (!r.success) {
      console.log('error getting title');
      return;
    }
    if (r.result.error) {
      linkErrCode = r.result.code;
      linkErr = true;
      linkTitle = '';
      linkTitleDisabled = false;
      return;
    }
    linkErr = false;
    linkTitle = r.result.title;
    linkTitleDisabled = false;
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

  async function create() {
    if (newTopics.length < 1 && type !== 'news') {
      console.log('at least one topic must be selected')
      return;
    }
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
      deleted: false,
      version: 1,
      last: true
    }

    if ([ 'task', 'note', 'link', 'news' ].includes(entry.type)) {
      entry.text = text;
    } else if (type === 'image') {
      // upload new images
      if (newImages.length > 0) {
        const res = await uploadNewImages();
        if (!res) {
          console.log('error at uploading images');
          return;
        }
      } else {
        console.log('no images for upload selected');
        return;
      }
      entry.images = newImages;
    }
    if (entry.type === 'link') {
      entry.comment = linkComment;
      entry.title = linkTitle;
    }

    const r = await sendRequest('POST', `/entry/${entry.id}.json`, entry);
    if (!r.success) {
      console.log('error creating entry');
      return;
    }

    console.log('success!')
    dispatch('created', entry);
    reset();
  }

  function reset() {
    text = '';
    newTopics = [];
    newTags = [];
    _private = false;
    pinned = false;
    linkComment = '';
    linkTitle = '';
    images = [];
    newImages = [];
    showAddElements = false;
    resetLoadImages = [];
  }

  function loadNewImages(images) {
    if (images.length > 0) {
      newImages = images;
      showAddElements = true;
    } else {
      reset();
    }
  }

</script>

<div class="newentry-box">
  <div>
    {#if type === 'image'}
      <LoadImages on:change={(e) => loadNewImages(e.detail)}
                  reset={resetLoadImages} />
    {:else if type === 'link'}
      <input placeholder="New Link..."
                bind:value={text}>
    {:else}
      <textarea class="newentry-text"
                placeholder="New Entry..."
                bind:value={text}></textarea>
    {/if}
  </div>
    {#if showAddElements}
      {#if type === 'link'}
        <div class="linktitle">
          <input id="linktitle" name="linktitle" placeholder="Link title..."
                 bind:value={linkTitle} disabled={linkTitleDisabled}>
          {#if linkErr}
            <small class="linkerr">Error getting title: {linkErrCode}</small>
          {/if}
        </div>
        <input id="linkcomment" name="linkcomment" placeholder="Link comment..."
               bind:value={linkComment}>
      {/if}
      {#if type !== 'news'}
        <EditTopics items={$currentTopics}
                    on:change={(e) => setNewTopics(e.detail)} />
        <EditTags tagsByTopics={$currentTagsByTopics} {newTopics}
                  on:change={(e) => setNewTags(e.detail)} />
      {/if}
    <div>
      {#if type !== 'news'}
        <input type="checkbox" id="private-checkbox" name="private"
               bind:checked={_private}>
        <label for="private-checkbox">Private</label>
      {/if}
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
    padding: 0 0 20px 0;
    border-bottom: 1px solid var(--main-line-color);
  }
  .newentry-text {
    width: 170px;
    height: 20px;
    padding: 10px;
  }
  .linktitle {
    display: flex;
    flex-flow: column;
  }
  .linkerr {
    color: var(--error-color);
  }
</style>
