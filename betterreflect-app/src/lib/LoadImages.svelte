<script>
  import { createEventDispatcher } from 'svelte';
  import { compressImage, encodeData } from '$lib/images';

  const dispatch = createEventDispatcher();

  export let reset = [];

  let images = [];
  let fileInputElement;

  $: _reset(reset);

  async function loadImages(f) {
    const n = await Promise.all(Array.from(f)
      .filter((file) => !images.map(v => v.filename).includes(file.name))
      .map(async (file) => {
        const blob = await compressImage(file, 240, 240);
        const data = await encodeData(blob);
        const image = {
          filename: file.name,
          osize: file.size,
          type: file.type,
          lastModified: file.lastModified,
          previewData: data,
          file: file,
          comment: "",
        };
        return image;
      })
    );
    images = [ ...images, ...n ];
    dispatch('change', images);
  }

  function unloadImage(filename) {
    images = images.filter(e => e.filename !== filename);
    dispatch('change', images);
  }

  function setComment(v, filename) {
    const image = images.find(e => e.filename === filename);
    image.comment = v;
    images = images;
    dispatch('change', images);
  }

  function _reset() {
    images = [];
    if (fileInputElement) fileInputElement.value = "";
  }
</script>

<input on:change={ (e) => loadImages(e.target.files) }
       bind:this={fileInputElement}
       type="file"
       accept="image/*"
       multiple>
{#if images.length > 0}
  <div class="images">
    {#each images as i}
      <div class="image">
        <img src={i.previewData} alt="loaded preview" />
        <button on:click={unloadImage(i.filename)}>Unload</button>
        <input bind:value={i.maxSize}
               title="Default: 1024px, larger images will be scaled to this size"
               placeholder="Max. image size..." />
        <textarea value={i.comment} placeholder="Add comment..."
               on:input={(e) => setComment(e.target.value, i.filename)} />
      </div>
    {/each}
  </div>
{/if}

<style>
  .images {
    padding-top: 10px;
    display: flex;
    gap: 15px;
  }
  .image {
    display: flex;
    flex-flow: column;
    align-items: flex-start;
    gap: 10px;
  }
</style>
