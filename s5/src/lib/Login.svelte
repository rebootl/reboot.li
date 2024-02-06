<script>

  /** @type {{
    *   clientData: import('$lib/types').ClientData,
    *   close: () => void
    * }} */
  let { clientData, close } = $props();

  // console.log('clientData', clientData);
</script>

<div class="user-menu-box">
  {#if clientData.loggedIn}
    <h2>{clientData.username}</h2>
<!--
    <div class="menuitem">
      <a href="/restore-delete"
         class:active={'/restore-delete' === $page.url.pathname}>Restore / Delete entries</a>
      <a href="/sessions" class:active={'/sessions' === $page.url.pathname}>Sessions</a>
    </div>
-->
    <form method="POST" action="/logout">
      <button>Logout</button>
    </form>
  {:else}
    <h2>Login</h2>
    <form method="POST" action="/login">
      <input name="username" placeholder="Username">
      <input name="password" type="password" placeholder="Password">
      <button>Login</button>
      <button type="button" class="secondary-button"
              onclick={close}
              aria-label="Close user menu"
              aria-controls="user-menu"
              >cancel</button>
    </form>
  {/if}
</div>

<style>
  .user-menu-box {
    display: flex;
    flex-direction: column;
    background-color: rgba(0, 0, 0, 0.9);
    border: 1px solid var(--primary-color-dimmed);
    border-radius: 15px;
    padding: 5px 25px 25px 25px;
  }
  form {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }
</style>
