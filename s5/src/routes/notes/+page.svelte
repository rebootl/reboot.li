<script>
	export let data;
  // console.log(data);
  
	let entries = data.entries;
  console.log(entries);

</script>

<h1>Notes</h1>

{#if data.clientData.loggedIn}
  <a href="/editNote/new">New Note...</a>
{/if}

<div class="note-list">
  {#each entries as entry}
    <div class="note-list-item">
      <a href="/note/{entry.id}">{entry.title}</a>
      <div>
        <small>{entry.created_at}</small>
        {#if entry.private}
          <small><span class="material-icons">lock</span> Private</small>
        {/if}
      </div>
      {#if data.clientData.loggedIn}
        <small><a href={ `/deleteNote/${entry.id}` }><span class="material-icons">delete</span></a></small>
      {:else}
        <span></span>
      {/if}
    </div>
  {:else}
    <p>No entries yet.</p>
  {/each}
</div>

<style>
  .note-list {
    margin-top: 30px;
    display: flex;
    flex-direction: column;
  }
  .note-list-item {
    display: flex;
    justify-content: space-between;
    border-bottom: 1px solid var(--primary-color-dimmed);
    padding-bottom: 10px;
    padding-top: 10px;
  }
  .note-list-item small {
    color: var(--text-color-dimmed);
  }
  small {
    font-size: 0.85em;
  }
  small .material-icons {
    font-size: 0.85em;
  }
</style>
