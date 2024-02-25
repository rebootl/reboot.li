<script>

  /** @typedef {Object} PropsData
   * @property {import('$lib/types').ClientData} clientData
   */
  /** @type {{ data: PropsData }} */
  let { data } = $props();

</script>

<div class="user-menu-box">
  <div class="login-box">
    {#if data.clientData.loggedIn}
      <h2>{data.clientData.username}</h2>
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
        <!--
        <button type="button" class="secondary-button"
                onclick={close}
                aria-label="Close user menu"
                aria-controls="user-menu"
                >cancel</button>-->
      </form>
    {/if}
  </div>
</div>

<style>
  .user-menu-box {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  form {
    display: flex;
    flex-direction: column;
    gap: 20px;
    max-width: 300px;
  }
</style>
