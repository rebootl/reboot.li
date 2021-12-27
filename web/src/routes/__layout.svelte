<svelte:head>
	<title>reboot.li - a personal website</title>
</svelte:head>

<script>
	import '../global.css';

	import { showMenu } from '$lib/store';
  import { page } from '$app/stores';

  import Login from '$lib/Login.svelte';
	import HeaderLinks from '$lib/HeaderLinks.svelte';

	let showLogin = false;

</script>

<div class="wrapper">
  <header class:fixed={$showMenu}>
    <div class="icon-button menu-button" on:click={() => $showMenu = !$showMenu}>
			<span class="material-icons">menu</span>
    </div>
    <div class="logo-box">
      <img class="logo" alt="Logo" src="/betterreflect-app-logo.png" />
    </div>

    <div class="right-box">
			<div class="header-links">
      	<HeaderLinks />
			</div>
      <div class="icon-button" on:click={() => showLogin = !showLogin}>
        <span class="material-icons">account_circle</span>
      </div>
    </div>
    <div class="login-box" class:show={showLogin}>
      <Login />
    </div>
  </header>

  <slot></slot>
</div>

<style>
  header {
    grid-area: header;
    display: flex;
    justify-content: space-between;
    background-color: var(--header-background-color);
  }
  .logo-box {
    display: flex;
    justify-content: center;
    align-items: center;
    width: var(--header-height);
    height: var(--header-height);
  }
  .logo {
    width: 38px;
    height: 38px;
  }
  .right-box {
    display: flex;
  }
  .icon-button {
    display: flex;
    justify-content: center;
    align-items: center;
    width: var(--header-height);
    height: var(--header-height);
  }
  .login-box {
    display: none;
    position: absolute;
    top: calc(var(--header-height) + 5px);
    right: 10px;
  }
  .login-box.show {
    display: flex;
  }
	.menu-button {
		display: none;
	}
	.header-links {
		display: flex;
	}
	@media all and (max-width: 600px) { /* 1000px = side width + max. main width */
		header.fixed {
		 	position: fixed;
			top: 0;
			left: 0;
			width: 100vw;
	 	}
		.menu-button {
			display: flex;
		}
		.header-links {
			display: none;
		}
	}
</style>
