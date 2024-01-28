<script>
  import dayjs from 'dayjs';

  import { goto } from '$app/navigation';
  import { sendRequest } from '$lib/request';

  // import LoadImages from '$lib/LoadImages.svelte';
  import { compressImage, encodeData } from '$lib/images';

  const manualDateFmt = 'YYYY-MM-DDTHH:mm';

	export let data;

  let images = [];

  // console.log(data);
	// let content = data.entry?.content;
  let manualDate = data.entry?.manual_date ? dayjs(data.entry?.manual_date).format(manualDateFmt) : '';
  // let title = data.entry?.title;

  async function confirmDelete() {
    if (!confirm("Are you sure you want to delete this entry?")) {
      return;
    }

    const r = await sendRequest('DELETE', `/entry/${data.entry?.id}`);
    if (!r.success) {
      console.log('error deleting entry');
      return;
    }

    console.log('success!')
    goto('/timeline');
  }

  async function loadImages(f) {
    console.log(f);
    const n = await Promise.all(Array.from(f)
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

  function resetImages() {
    images = [];
    const fileInputElement = document.querySelector('#images-file-input');
    fileInputElement.value = "";
  }

  /*function unloadImage(filename) {
    images = images.filter(e => e.filename !== filename);
  }*/

  /*function setComment(v, filename) {
    const image = images.find(e => e.filename === filename);
    image.comment = v;
    images = images;
  }*/

  async function submit() {

    // compress images
    const maxSizeEl = document.querySelector('#max-size');
    const maxSize = maxSizeEl ? parseInt(maxSizeEl.value) : 1024;
    console.log(maxSize);

    const formData = new FormData();
    // -> use promise.all
    for (const image of images) {
      const blob = await compressImage(image.file, maxSize, maxSize);
      formData.append('images', blob, image.filename);
    }

    // add other form data
    const form = document.querySelector('#new-entry-form');
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
    console.log(response);
    if (!response) {
      console.log('error creating entry');
      return;
    }

    // reset file input element
    // const fileInputElement = document.querySelector('#images-file-input');
    // fileInputElement.value = "";

    // const form = document.querySelector('#new-entry-form');
    // form.submit();
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
      id="images-file-input"/>
    {#if images.length > 0}
      <div>
        <button type="button" class="small-button" onclick={() => resetImages()}>Reset images</button>
        <label>
          Max size:
          <input type="number" id="max-size" value="1024" />
        </label>
      </div>
    {/if}
    {#each images as image}
      <div class="image-load-preview">
        <img src={image.previewData} alt={image.filename} width="240" height="240" />
        <input type="text" name="imagecomment" placeholder="Comment" />
        <!--<button type="button" onclick={() => unloadImage(image.filename)}>Unload</button>-->
      </div>
    {/each}
    <div>
      <label>
        <input type="checkbox" name="isPrivate" />
        Private
      </label>
    </div>
    <div class="buttons">
      <button type="button" onclick={() => submit()}>Create</button>
      <a href="/timeline">Cancel</a>
    </div>
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
</style>
