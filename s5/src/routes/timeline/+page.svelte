<script>
  import { browser } from '$app/environment';

  import ImageViewer from '$lib/ImageViewer.svelte';

  /** @typedef {Object} Data
    * @property {import('$lib/types').TimelineEntry[]} timelineEntries
    * @property {import('$lib/types').ClientData} clientData
    */
  /** @type {{ data: Data }} */
  let { data } = $props();

  let showImageViewer = $state(false);
  /** @type {import('$lib/server/db.js').ImageData[]} */
  let imageViewerImages = $state([]);
  /** @type {number} */
  let imageViewerCurrentImageIdx = $state(0);

  /** @type {HTMLElement|null} */
  let openImageViewerElement = null;

  /** @param {KeyboardEvent} e */
  function keyPressHandler(e) {
    if (e.key === 'Escape') {
      closeImageViewer();
    }
  }

  if (browser) {
    window.addEventListener('keydown', keyPressHandler);
  };

  /** @param {KeyboardEvent} e
    * @param {import('$lib/server/db.js').EntryData } entry
    * @param {number} i */
  function handleImageKeyPress(e, entry, i) {
    if (e.key === 'Enter') {
      openImageViewer(e, entry, i);
    }
  }

  /** @param {import('$lib/server/db.js').EntryData } entry
    * @param {number} i */
  function openImageViewer(e, entry, i) {
    openImageViewerElement = e.target;
    
    imageViewerImages = entry.images;
    imageViewerCurrentImageIdx = i;
    showImageViewer = true;
  }

  function closeImageViewer() {
    // console.log('openImageViewerElement', openImageViewerElement);
    showImageViewer = false;
    // openImageViewerElement?.focus();
  }
</script>

<h1>Timeline</h1>

{#if data.clientData.loggedIn}
  <a href="/editEvent/new">New Event...</a>
{/if}

{#if showImageViewer}
  <ImageViewer images={imageViewerImages} currentImageIdx={imageViewerCurrentImageIdx} show={showImageViewer}
    close={() => closeImageViewer()}/>
{/if}

<div class="list">
  {#each data.timelineEntries as t}
    {#if t.type === 'year'}
      <h2 class="year">{t.year}</h2>
    {:else if t.type === 'month'}
      <h3 class="month">{t.month}</h3>
    {:else if t.type === 'entry'}
      <div class="list-item">
        <div class="item-header">
          <div>
            <small>{ t.date }</small>
            {#if t.entry?.private}
              <small><span class="material-icons">lock</span> Private</small>
            {/if}
          </div>
          {#if data.clientData.loggedIn}
            <small><a href={ `/editEvent/${t.entry?.id}` }><span class="material-icons">edit</span></a></small>
          {:else}
            <span></span>
          {/if}
        </div>
        <div class="item-content">
          {t.entry?.content}
          {#if t.entry?.comment}
            <small>{t.entry.comment}</small>
          {/if}
          {#if t.entry?.images}
            <div class="images-preview-box">
              {#each t.entry.images as image, i}
                {#if browser}
                  <div class="image-preview-box" tabindex="0"
                       onclick={ (e) => openImageViewer(e, t.entry, i) }
                       onkeyup={ (e) => handleImageKeyPress(e, t.entry, i) }
                       role="button"
                       aria-label="Show image in overlay"
                       aria-controls="image-viewer"
                       aria-expanded={showImageViewer}
                       aria-haspopup="dialog"
                       >
                    <img class="image-preview" alt={ image.comment } src={ 'data:image/png;base64,' + image.preview_data } />
                  </div>
                {:else}
                  <a href={ image.path } target="_blank">
                    <img class="image-preview" alt={ image.comment } src={ 'data:image/png;base64,' + image.preview_data } />
                  </a>
                {/if}
              {/each}
            </div>
          {/if}
        </div>
      </div>
    {/if}
  {:else}
    <p>No entries yet.</p>
  {/each}
</div>

<style>
  form {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }
  .list {
    margin-top: 30px;
    margin-bottom: 30px;
    display: flex;
    flex-direction: column;
    gap: 15px;
  }
  .list-item {
    display: flex;
    flex-direction: column;
    border-radius: 15px;
  }
  .list-item .item-header {
    display: flex;
    justify-content: space-between;
    color: var(--text-color-dimmed);
    font-size: 0.85em;
  }
  .list-item .item-content {
    display: flex;
    flex-direction: column;
    gap: 5px;
    margin-left: 20px;
    padding: 20px 40px 20px 40px;
    border-left: 1px solid var(--primary-color-dimmed);
  }
  small .material-icons {
    font-size: 0.85em;
  }
  .year {
    margin-top: 0.5em;
    margin-bottom: 0;
    color: var(--text-color-dimmed);
  }
  .month {
    color: var(--text-color-dimmed);
  }
  .images-preview-box {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
  }
  .image-preview-box:focus {
    outline: 2px solid var(--secondary-color);
  }
  .image-preview {
    max-width: 120px;
    max-height: 120px;
    object-fit: contain;
  }
</style>
