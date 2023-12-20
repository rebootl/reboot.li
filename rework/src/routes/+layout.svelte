<svelte:head>
	<title>reboot.li - my personal website</title>
</svelte:head>

<script>
	import '../global.css';

  import 'material-icons/iconfont/material-icons.css';
  // import Login from '$lib/Login.svelte';
	// import HeaderLinks from '$lib/HeaderLinks.svelte';
  let showMenu = false;
  
</script>
<!--
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
-->

<header>
  <div class="icon-button" on:click={() => showMenu = !showMenu}>
		<span class="material-icons">menu</span>
  </div>
  <div class="logo-box">
    <img class="logo" alt="Logo" src="/betterreflect-app-logo.png" />
  </div>
  <div class="icon-button">
    <span class="material-icons">account_circle</span>
  </div>
</header>
<div class="wrapper">
  <aside class:shown={showMenu}>
  </aside>
  <div class="main-wrapper" class:menu-shown={showMenu}>
    <main>
      <slot></slot>
    </main>
  </div>
</div>

<style>
  header {
    height: var(--header-height);
    border-bottom: 2px solid var(--primary-color);
    display: flex;
    justify-content: space-between;
    position: fixed;
    width: 100%;
    top: 0;
    left: 0;
    z-index: 100;
    background-color: var(--background-color);
    /* background-color: var(--header-background-color); */
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
  /*.right-box {
    display: flex;
  }*/
  .icon-button {
    display: flex;
    justify-content: center;
    align-items: center;
    width: var(--header-height);
    height: var(--header-height);
    cursor: pointer;
  }
  /*.login-box {
    display: none;
    position: absolute;
    top: calc(var(--header-height) + 5px);
    right: 10px;
  }
  .login-box.show {
    display: flex;
  }*/
	/*.menu-button {
		display: flex;
	}*/
	/*.header-links {
		display: flex;
	}*/
	/*@media all and (max-width: 600px) { /* 1000px = side width + max. main width */
	/*	header.fixed {
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
	}*/
  .wrapper {
    margin-top: var(--header-height);
    position: relative;
  }
  .main-wrapper {
    display: flex;
    justify-content: center;
    position: absolute;
    top: 0;
    left: 0;
    width: calc(100%);
    min-width: var(--min-main-width);
    transition: left 0.2s ease-in-out, width 0.2s ease-in-out;
  }
  .main-wrapper.menu-shown {
    left: var(--side-width);
    width: calc(100% - var(--side-width));
  }
  main {
    max-width: var(--max-main-width);
  }
  aside {
    margin-top: var(--header-height);
    width: var(--side-width);
    border: 1px solid var(--primary-color);
    position: fixed;
    top: 0;
    left: calc(var(--side-width) * -1);
    height: calc(100vh - var(--header-height));
    transition: left 0.2s ease-in-out;
  }
  aside.shown {
    left: 0;
  }
</style>
