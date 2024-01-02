<script>
	export let data;
  // console.log(data);
</script>

{#if data.entry}
  <h1>Edit Note</h1>
  <form method="POST" action={ `/edit_note/${ data.entry.id }?/updateEntry` }>
    <input type="text" name="title" placeholder="Title" value={ data.entry.title } />
    <textarea name="content" placeholder="Text...">{ data.entry.content }</textarea>
    <div>
      <label>
        <input type="checkbox" name="isPrivate" checked={ data.entry.private ? true : false }/>
        Private
      </label>
    </div>
    <div class="buttons">
      <button>Save</button>
      <a href="/notes">Cancel</a>
    </div>
  </form>
{:else}
  <h1>New Note</h1>
  <form method="POST" action="/edit_note/new?/createEntry">
    <input type="text" name="title" placeholder="Title" />
    <textarea name="content" placeholder="Text..."></textarea>
    <div>
      <label>
        <input type="checkbox" name="isPrivate" />
        Private
      </label>
    </div>
    <div class="buttons">
      <button>Create</button>
      <a href="/notes">Cancel</a>
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
  }
</style>
