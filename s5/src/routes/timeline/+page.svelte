<script>
	export let data;
  // console.log(data);
  
	let entries = data.entries;
  console.log(entries);

</script>

<h1>Timeline</h1>

{#if data.clientData.loggedIn}
  <a href="/editPost/new">New Post...</a>
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
          <small><a href={ `/editPost/${entry.id}` }><span class="material-icons">edit</span></a></small>
        {:else}
          <span></span>
        {/if}
      </div>
      <div class="item-content">
        {entry.content}>
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
    background-color: #1e1e1e;
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
