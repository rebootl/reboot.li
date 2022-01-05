<script context="module">
  export async function load({ page, fetch, session, context }) {

  	const url = '/restore-delete.json';

  	const res = await fetch(url);
  	if (res.ok) {
  		return {
  			props: {
  				entries: await res.json(),
  			}
  		};
  	}

  	return {
  		status: res.status,
  		error: new Error(`Could not load ${url}`)
  	};
  }
</script>

<script>
  import moment from 'moment';
  import SideNav from '$lib/SideNav.svelte';
  import { sendRequest, getToken } from '$lib/request';

  const dateFormat = 'MMM D YYYY - HH:mm';

  export let entries = [];

  async function restore(entry) {
    if (!confirm("Restore entry?"))
      return;
    const r = await sendRequest('PUT', '/restore-delete.json', {
      id: entry.id,
      user: entry.user
    });
    if (!r.success) {
      console.log('error restoring entry');
      return;
    }
    entries = entries.filter(e => e.id !== entry.id);
    console.log('success!')
  }

  async function deletePermanently(entry) {
    if (!confirm("This entry (incl. images) will be permanently deleted from the server!"))
      return;

    if (entry.type === 'image') {
      // get mediaserver token
      const token = await getToken();
      if (!token) {
        console.log('error getting mediaserver token');
        return;
      }

      for (const i of entry.images) {
        const r = await sendTokenRequest('POST',
          new URL('/api/deleteImage', MEDIASERVER),
          { filepath: i.filepath },
          token
        );
        if (!r.success) {
          console.log('error deleting image');
          return;
        }
      }
    }

    const r = await sendRequest('DELETE', '/restore-delete.json', {
      id: entry.id,
      user: entry.user
    });
    if (!r.success) {
      console.log('error deleting entry');
      return;
    }
    entries = entries.filter(e => e.id !== entry.id);
    console.log('success!')
  }

</script>

<SideNav showOnWide={false} />
<main>
  <h1>Restore / Delete entries</h1>
  {#each entries as e}
    <div class="box">
      <div class="smallinfo">
        <small>Created: {moment(new Date(e.date)).format(dateFormat)}</small>
      </div>
      <div class="maininfo">
        <div class="maininfo-left">
          <div class="maininfo-title">
            <span class="material-icons">link</span>
            {e.id}
          </div>
          <pre>{e.text}</pre>
        </div>
        <div class="buttons-right">
          <button on:click={() => restore(e)}>
            Restore
          </button>
          <button on:click={() => deletePermanently(e)} class="deletebutton">
            Delete permanently
          </button>
        </div>
      </div>
      <div class="smallinfo">
        <small>Deleted: {moment(new Date(e.deleteDate)).format(dateFormat)}</small>
      </div>
    </div>
  {:else}
    <small class="info">
      loading...
    </small>
  {/each}
</main>

<style>
  .box {
    border-bottom: 1px solid var(--main-line-color);
    max-height: 400px;
    overflow: auto;
  }
  .maininfo {
    display: flex;
    justify-content: space-between;
    padding-top: 5px;
    padding-bottom: 5px;
  }
  .maininfo-left {
    display: flex;
    flex-flow: column;
  }
  .maininfo-title {
    display: flex;
    gap: 10px;
  }
  pre {
    margin-bottom: 0;
  }
  .smallinfo {
    display: flex;
    gap: 10px;
    color: var(--main-text-label-color);
    font-size: 0.85em;
  }
  .buttons-right {
    display: flex;
    flex-flow: column;
    justify-content: space-between;
    align-items: flex-end;
  }
  .deletebutton {
    background-color: var(--error-color);
    color: var(--on-error-color);
    border: 2px solid var(--error-color);
    border-radius: 5px;
  }
</style>
