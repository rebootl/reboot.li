<script>
  import dayjs from 'dayjs';

  import { goto } from '$app/navigation';
  import { sendJSONRequest } from '$lib/request';

  import { compressImage, encodeData } from '$lib/images';

  const manualDateFmt = 'YYYY-MM-DDTHH:mm';

  // -> these types should be imported from db!!
  /** @typedef {Object} ImageData
   * @property {File} file
   * @property {string} path
   * @property {string} preview_data
   * @property {string} [comment]
   */
  /** @typedef {Object} EntryData
   * @property {number} id
   * @property {string} content
   * @property {string} comment
   * @property {string} date
   * @property {string} manual_date
   * @property {boolean} private
   * @property {ImageData[]} images
   */
  /** @typedef {Object} Data
   * @property {EntryData} entry
   */
  /** @type {Data} */
	export let data;

  /** @typedef {Object} CreateImageData
   * @property {File} file
   * @property {string} filename
   * @property {string} previewData
   * @property {string} [comment]
   */
  /** image array used for preview
    * @type {CreateImageData[]}
    */
  let images = [];

  /** max image size
    * @type {string}
    */
  let maxImageSize = '1024';

  /** override date with manual date if set
    * @type {string}
    */
  let manualDate = data.entry?.manual_date ? dayjs(data.entry?.manual_date).format(manualDateFmt) : '';

  /** when deleting an entry we want to at least show a confirmation dialog,
    * for now we use this function to do that
    *
    * we use a JSON request here because the button is located within the form
    * to save the edit, but we do not want to submit this form
    *
    * @todo it would be nice to find a better way to do this, maybe simply a separate
    *       form for that is triggered by this button
    */
  async function confirmDelete() {
    if (!confirm("Are you sure you want to delete this entry?")) {
      return;
    }

    const r = await sendJSONRequest('DELETE', `/entry/${data.entry?.id}`);
    if (!r.success) {
      console.log('error deleting entry');
      return;
    }

    console.log('success!')
    goto('/timeline');
  }

  /** when images are selected we want to show a preview and a comment input,
    * we generate the preview here and add the images to the images array
    * @param {FileList} files
    */
  async function loadImages(files) {
    // console.log(files);
    const n = await Promise.all(Array.from(files)
      .map(async (file) => {
        const blob = await compressImage(file, 240, 240);
        const data = await encodeData(blob);
        const image = {
          file: file,
          filename: file.name,
          previewData: data,
        };
        return image;
      })
    );
    images = n;
  }

  /** reset image file selection
    */
  function resetImages() {
    images = [];
    const fileInputElement = /** @type {HTMLInputElement} */ (document.querySelector('#images-fie-input'));
    fileInputElement.value = "";
  }

  /** create entry via function in order to pre-compress the images
    */
  async function createEntry() {

    // console.log('maxImageSize', maxImageSize);
    const maxSize = !isNaN(parseInt(maxImageSize)) ? parseInt(maxImageSize) : 1024;

    // create a new form for submission
    const formData = new FormData();

    // compress images
    // time it
    // console.time('compress images');
    await Promise.all(images.map(async (image) => {
      const blob = await compressImage(image.file, maxSize, maxSize);
      formData.append('images', blob, image.filename);
    }));
    // for (const image of images) {
    //   const blob = await compressImage(image.file, maxSize, maxSize);
    //   formData.append('images', blob, image.filename);
    // }
    // console.timeEnd('compress images');

    // add other form data
    const form = /** @type {HTMLFormElement} */ (document.querySelector('#new-entry-form'));
    const formDataOrig = new FormData(form);
    for (const [key, value] of formDataOrig.entries()) {
      if (key === 'images') continue;
      formData.append(key, value);
    }

    // send request
    const response = await fetch('/editPost/new?/createEntry', {
      method: "POST",
      body: formData,
    });
    // console.log(response);
    if (!response) {
      console.log('error creating entry');
      return;
    }

    console.log('success!')
    // reset images
    images = [];
    goto('/timeline');
  }

</script>

{#if data.entry}
  <h1>Edit Post</h1>
  <form method="POST" action={ `/editPost/${ data.entry?.id }?/updateEntry` }>
    <textarea name="content" placeholder="Text...">{ data.entry.content }</textarea>
    <input type="text" name="comment" placeholder="Comment" value={ data.entry?.comment } />
    <label>
      <input type="datetime-local" name="manualdate" value={ manualDate } />
      Set to use manual date
    </label>
    <div>
      <label>
        <input type="checkbox" name="isPrivate" checked={ data.entry?.private ? true : false }/>
        Private
      </label>
    </div>
    <div class="buttons">
      <div>
        <button>Save</button>
        <a href="/timeline">Cancel</a>
      </div>
      <button type="button" class="danger-button" onclick={() => confirmDelete()}>Delete</button>
    </div>
  </form>
  <h2>Edit Images</h2>
  {#each data.entry.images as image}
    <div class="image-load-preview">
      <img src={'data:image/png;base64,' + image.preview_data} alt={image.path} class="image-preview"
        title={image.path} />
      <input type="text" name="imagecomment" placeholder="Comment" value={image.comment} />
      <div>
        <button class="danger-button small-button" onclick={() => deleteImage(image.path)}>Delete image</button>
      </div>
    </div>
  {/each}
{:else}
  <h1>New Post</h1>
  <form method="POST" action="/editPost/new?/createEntry" enctype="multipart/form-data" id="new-entry-form">
    <textarea name="content" placeholder="Text..."></textarea>
    <input type="text" name="comment" placeholder="Comment" />
    <label>
      <input type="datetime-local" name="manualdate" value={ data.entry?.date } 
              on:change={(e) => loadImages(e.target.files) } /> 
    </label>
    <input type="file" name="images" accept="image/*" multiple onchange={(e) => loadImages(e.target.files) }
      id="images-file-input" />
    {#if images.length > 0}
      <div>
        <button type="button" class="small-button" onclick={() => resetImages()}>Reset images</button>
      </div>
    {/if}
    <div>
      <label>
        <input type="checkbox" name="isPrivate" />
        Private
      </label>
    </div>
    <div class="buttons">
      <button type="button" onclick={() => createEntry()}>Create</button>
      <a href="/timeline">Cancel</a>
    </div>
    {#if images.length > 0}
      <h2>Images</h2>
      <div>
        <label>
          Max. image size:
          <input type="text" id="max-size" bind:value={maxImageSize} />
        </label>
      </div>
    {/if}
    {#each images as image}
      <div class="image-load-preview">
        <img src={image.previewData} alt={image.filename} class="image-preview" title={image.filename} />
        <input type="text" name="imagecomment" placeholder="Comment" />
        <!--<button type="button" onclick={() => unloadImage(image.filename)}>Unload</button>-->
      </div>
    {/each}
  </form>
{/if}

<style>
  form {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }
  textarea {
    height: 160px;
  }
  .buttons {
    display: flex;
    gap: 20px;
    align-items: center;
    justify-content: space-between;
  }
	.small-button {
		display: flex;
	  align-items: center;
  }
  .image-load-preview {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }
  .image-preview {
    max-width: 240px;
    max-height: 240px;
    object-fit: contain;
  }
</style>
