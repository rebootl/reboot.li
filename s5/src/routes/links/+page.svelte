<script>
  import { debounce } from '$lib/helper';
  import { sendJSONRequest } from '$lib/request';

	export let data;
  // console.log(data);
  
	let entries = data.entries;
  console.log(entries);

  let showAddLink = false;
  let title = '';
  let titleDisabled = true;
  let linkErr = false;
  let linkErrMessage = '';

  /** @param {Event} event */
  function linkInsert(event) {
    const v = /** @type {HTMLInputElement} */ (event.target).value;
  
    if (v === "") {
      showAddLink = false;
      return;
    }
    showAddLink = true;

    titleDisabled = true;
    title = 'getting title...';
    debounce(() => getTitle(v), 500);
  }

  /** @param {string} text */
  async function getTitle(text) {
    console.log('getTitle')
    const r = await sendJSONRequest('POST', '/getTitle', {
      url: text
    });
    if (!r.success) {
      console.log('error getting title');
      return;
    }
    if (r.result.error) {
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
</script>

<h1>Collected Links</h1>

{#if data.clientData.loggedIn}
  <form action="/editLink/new?/createEntry" method="POST">
    <input type="text" name="url" placeholder="Add link" oninput={ (ev) => linkInsert(ev) } />
    {#if showAddLink}
      <input type="text" name="title" placeholder="Link title..."
             bind:value={title} disabled={titleDisabled}>
      {#if linkErr}
        <small class="warning">Error getting title: {linkErrMessage}</small>
      {/if}
      <input type="text" name="comment" placeholder="Comment" />
      <button>Add</button>
    {/if}
  </form>
{/if}

<div class="list">
  {#each entries as entry}
    <div class="list-item">
      <div class="item-header">
        <small>{entry.created_at}</small>
        {#if entry.private}
          <small><span class="material-icons">lock</span> Private</small>
        {/if}
        {#if data.clientData.loggedIn}
          <small><a href={ `/editLink/${entry.id}` }><span class="material-icons">edit</span></a></small>
        {:else}
          <span></span>
        {/if}
      </div>
      <div class="item-content">
        <a class="item-link" href={entry.content}><small>{entry.content}</small></a>
        {entry.title}
        {#if entry.comment}
          <small>{entry.comment}</small>
        {/if}
      </div>
    </div>
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
    display: flex;
    flex-direction: column;
    gap: 15px;
  }
  .list-item {
    display: flex;
    flex-direction: column;
    padding: 20px;
    background-color: var(--card-background-color);
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
  }
  small .material-icons {
    font-size: 0.85em;
  }
</style>
