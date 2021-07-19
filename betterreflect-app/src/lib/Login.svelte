<script>
  import { session } from '$app/stores';

  let username = '';
  let password = '';

  let loggedIn = $session.loggedIn;
  //console.log(session)
  //console.log(loggedIn)

  async function login() {
    try {
      const res = await fetch('/login.json', {
        method: 'POST',
        body: JSON.stringify({
          user: username,
          password: password
        })
      });
			if (res.ok) {
        // -> what to do
        console.log('login successful')
        username = '';
        password = '';
        loggedIn = true;
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
      const res = await fetch('/logout.json', {
        method: 'POST',
      });
			if (res.ok) {
        console.log('logout successful')
        loggedIn = false;
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
    <h2>Logout</h2>

    <button on:click={logout}>Logout</button>
  {:else}
    <h2>Login</h2>

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
    /*border: 1px solid var(--primary-color);*/
    border-radius: 5px;
  }
  input {
    width: 160px;
    margin-bottom: 15px;
  }
  button {
    margin-top: 5px;
  }
</style>
