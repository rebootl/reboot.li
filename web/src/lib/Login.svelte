<script>
  import { page } from '$app/stores';

  let username = '';
  let password = '';

  let loggedIn = $page.data.loggedIn;

  async function login() {
    try {
      const res = await fetch('/login', {
        method: 'POST',
        body: JSON.stringify({
          user: username,
          password: password
        })
      });
			if (res.ok) {
        console.log('login successful')
        username = '';
        password = '';
        window.location.reload();
				return;
			}
			const { message } = await res.json();
	    console.warn(message);
		} catch(error) {
			console.error(error);
		}
  }

  async function logout() {
    try {
      const res = await fetch('/logout', {
        method: 'POST',
      });
			if (res.ok) {
        console.log('logout successful')
        // reload page
        window.location.reload();
				return;
			}
			const { message } = await res.json();
	    console.warn(message);
		} catch(error) {
			console.error(error);
		}
  }

</script>

<div class="login-box">
  {#if loggedIn}
    <h2>{$page.data.user}</h2>

    <div class="menuitem">
      <a href="/restore-delete"
         class:active={'/restore-delete' === $page.url.pathname}>Restore / Delete entries</a>
      <a href="/sessions" class:active={'/sessions' === $page.url.pathname}>Sessions</a>
    </div>

    <button on:click={logout}>Logout</button>
  {:else}
    <h2>Login</h2>
    <!--<form method="POST">
      <input name="username" placeholder="Username">
      <input name="password" type="password" placeholder="Password">
      <button>Log in</button>
    </form>-->
    <input bind:value={username} placeholder="Username">
    <input type="password" bind:value={password} placeholder="Password">

    <button on:click={login}>Login</button>
  {/if}
</div>

<style>
  .login-box {
    display: flex;
    flex-direction: column;
    padding: 0 20px 20px 25px;
    background-color: var(--dialog-background-color);
    border-radius: 5px;
  }
  input {
    width: 160px;
    margin-bottom: 15px;
  }
  h2 {
    margin-top: revert;
  }
  .menuitem {
    margin-bottom: 20px;
    display: flex;
    flex-flow: column;
    gap: 20px;
  }
  .menuitem a {
    color: var(--menu-link-color);
    text-decoration: none;
  }
  .menuitem a.active {
    text-decoration: underline;
  }
  button {
    margin-top: 5px;
  }
</style>
