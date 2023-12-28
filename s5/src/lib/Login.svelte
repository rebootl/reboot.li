<script>
  
  /** @type {import('$lib/types').clientData} */
  export let clientData;

  // console.log('clientData', clientData);
</script>

<div class="login-box">
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
