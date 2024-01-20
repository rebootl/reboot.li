<script>
  import { goto } from '$app/navigation';
  import { sendRequest } from '$lib/request';

	export let data;
  // console.log(data);
	let content = data.entry?.content;
  let title = data.entry?.title;
  let titleDisabled = true;
  let linkErr = false;
  let linkErrMessage = '';

  async function getTitle() {
    console.log('getTitle')
    const r = await sendRequest('POST', '/getTitle', {
      url: content
    });
    if (!r.success) {
      console.log('error getting title');
      return;
    }
    if (r.result.error) {
			console.log(r.result);
      
      linkErrMessage = r.result.message;
      linkErr = true;
      title = '';
      titleDisabled = false;
      return;
    }
    linkErr = false;
    title = r.result.title;
    titleDisabled = false;
  }

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
    goto('/links');
  }
</script>

<h1>Edit Link</h1>
<form method="POST" action={ `/edit_link/${ data.entry?.id }?/updateEntry` }>
  <input type="text" name="content" placeholder="URL" bind:value={content} />
  <input type="text" name="title" placeholder="Title" bind:value={title} disabled={titleDisabled} />
	<div>
	  <button type="button" class="small-button" on:click={() => getTitle()}>
			<span class="material-icons">refresh</span> Update title
		</button>
		{#if linkErr}
		  <small class="warning">Error: {linkErrMessage}</small>
		{/if}
  </div>
  <input type="text" name="comment" placeholder="Comment" value={ data.entry?.comment } />
  <div>
    <label>
      <input type="checkbox" name="isPrivate" checked={ data.entry?.private ? true : false }/>
      Private
    </label>
  </div>
  <div class="buttons">
    <div>
      <button>Save</button>
      <a href="/links">Cancel</a>
    </div>
    <button type="button" class="danger-button" onclick={() => confirmDelete()}>Delete</button>
  </div>
</form>

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
</style>
