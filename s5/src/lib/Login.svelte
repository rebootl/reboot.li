<script>
  /**
   * @type {boolean}
   */
  let loggedIn = false;
</script>

<div class="login-box">
  {#if loggedIn}
    <h2>[username]</h2>
<!--
    <div class="menuitem">
      <a href="/restore-delete"
         class:active={'/restore-delete' === $page.url.pathname}>Restore / Delete entries</a>
      <a href="/sessions" class:active={'/sessions' === $page.url.pathname}>Sessions</a>
    </div>
-->
    <button on:click={_=>{}}>Logout</button>
  {:else}
    <h2>Login</h2>
    <form method="POST">
      <input name="username" placeholder="Username">
      <input name="password" type="password" placeholder="Password">
      <button>Login</button>
    </form>
  {/if}
</div>

<style>
  .login-box {
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
