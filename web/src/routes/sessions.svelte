<script context="module">
  export async function load({ page, fetch, session, context }) {

  	const url = '/sessions.json';

  	const res = await fetch(url);
  	if (res.ok) {
  		return {
  			props: {
  				sessions: await res.json(),
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
  import { session } from '$app/stores';
  import { sendRequest } from '$lib/request';

  export let sessions = [];

  async function revoke(_id) {
    if (!confirm("Do u really want to delete this session?"))
      return;

    const r = await sendRequest('DELETE', '/sessions.json', {
      user: $session.user,
      _id: _id
    });
    if (!r.success) {
      console.log('error deleting entry');
      return;
    }
    sessions = sessions.filter(s => s._id !== _id);
    console.log('success!')
  }

</script>

<SideNav showOnWide={false} />
<main>
  <h1>Sessions</h1>
  {#each sessions as s}
    <div class="session-box">
      <div class="smallinfo">
        <small>Created {moment(new Date(s.createdAt)).fromNow()}</small>
        {#if s.current}
          <small>[ Current ]</small>
        {/if}
      </div>
      <div class="maininfo">
        <div class="maininfo-left">
          <span class="material-icons">computer</span>
          {s.host}
        </div>
        {#if !s.current}
          <button on:click={() => revoke(s._id)} class="deletebutton">
            Revoke
          </button>
        {/if}
      </div>
      <div class="smallinfo">
        <small>{s.userAgent}</small>
      </div>
    </div>
  {:else}
    <small class="info">
      loading...
    </small>
  {/each}
</main>

<style>
  .session-box {
    border-bottom: 1px solid var(--main-line-color);
  }
  .maininfo {
    display: flex;
    justify-content: space-between;
    padding-top: 5px;
    padding-bottom: 5px;
  }
  .maininfo-left {
    display: flex;
    gap: 10px;
  }
  .smallinfo {
    display: flex;
    gap: 10px;
    color: var(--main-text-label-color);
    font-size: 0.85em;
  }
  .deletebutton {
    background-color: var(--error-color);
    color: var(--on-error-color);
    border: 2px solid var(--error-color);
    border-radius: 5px;
  }
</style>
