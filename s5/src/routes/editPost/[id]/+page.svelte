<script>
  import dayjs from 'dayjs';

  import { goto } from '$app/navigation';
  import { sendRequest } from '$lib/request';

  const manualDateFmt = 'YYYY-MM-DDTHH:mm';

	export let data;
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
  <form method="POST" action="/editPost/new?/createEntry">
    <textarea name="content" placeholder="Text..."></textarea>
    <input type="text" name="comment" placeholder="Comment" />
    <label>
      <input type="datetime-local" name="manualdate" value={ data.entry?.date } />
      Set to use manual date
    </label>
    <div>
      <label>
        <input type="checkbox" name="isPrivate" />
        Private
      </label>
    </div>
    <div class="buttons">
      <button>Create</button>
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
</style>
