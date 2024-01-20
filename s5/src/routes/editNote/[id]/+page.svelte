<script>
  import { goto } from '$app/navigation';
  import { sendRequest } from '$lib/request';

	export let data;
  // console.log(data);

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
    goto('/notes');
  }
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
      <div>
        <button>Save</button>
        <a href="/notes">Cancel</a>
      </div>
      <button type="button" class="danger-button" onclick={() => confirmDelete()}>Delete</button>
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
    justify-content: space-between;
  }
</style>
