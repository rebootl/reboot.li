<svelte:head>
	<title>reboot.li - my personal website</title>
</svelte:head>

<script>
  import { browser } from '$app/environment';

  import 'material-icons/iconfont/material-icons.css';
	import '../global.css';
  import Login from '$lib/Login.svelte';

  /** @typedef {Object} PropsData
   * @property {import('$lib/types').ClientData} clientData
   */
  /** @type {{ data: PropsData }} */
  let { data } = $props();
  // console.log('data layout', data);

  /** @type {boolean} */
  let showMenu = $state(false);

  /** @type {boolean} */
  let showLogin = $state(false);

  if (browser) {
    document.addEventListener('keydown', (e) => {
      if (e.key === 'Escape') {
        showMenu = false;
        showLogin = false;
      }
    });
  }
</script>

<header>
  <button id="main-menu-button" class="icon-button"
          onclick={() => showMenu = !showMenu}
          aria-haspopup="true"
          aria-expanded={showMenu}
          aria-controls="main-menu"
          aria-label="Open / close main menu"
          >
		<span class="material-icons">menu</span>
  </button>
  <div class="logo-box">
    <img class="logo" alt="" src="/logo.png" />
  </div>
  <button id="user-menu-button" class="icon-button"
          onclick={() => showLogin = !showLogin}
          aria-haspopup="true"
          aria-expanded={showLogin}
          aria-controls="user-menu"
          aria-label="Open / close user menu"
          >
    <span class="material-icons">account_circle</span>
  </button>
</header>
<div class="wrapper">
  <aside id="main-menu" class:shown={showMenu}
          role="navigation"
          aria-label="Main menu"
          >
    <nav>
      <ul>
        <li><a href="/" onclick={() => showMenu = false}>Home</a> | </li>
        <li><a href="/timeline" onclick={() => showMenu = false}>Timeline</a> | </li>
        <li><a href="/notes" onclick={() => showMenu = false}>Notes</a> | </li>
        <li><a href="/links" onclick={() => showMenu = false}>Links</a></li>
      </ul>
      <button class="secondary-button"
              onclick={() => showMenu = false}
              aria-label="Close main menu"
              aria-controls="main-menu"
              >Close menu</button>
    </nav>
    <footer class="credits">
      <small>Copyright 2024 Cem Aydin<br />
        Created with <a href="https://kit.svelte.dev/" target="_blank">SvelteKit</a></small>
    </footer>
  </aside>
  <div id="user-menu" class="login-box" class:shown={showLogin}
       role="menu"
       aria-label="User menu"
       >
    <Login clientData={data.clientData} close={() => showLogin = false} />
  </div>
  <div class="overlay" class:shown={showMenu || showLogin}
       onclick={() => {showMenu = false ; showLogin = false}}
       onkeydown={(e) => {if (e.key === 'Enter' || e.key === 'Space') {showMenu = false ; showLogin = false}}}
       role="button"
       aria-label="Close menu"
       aria-hidden={!showMenu && !showLogin}
       aria-controls="main-menu user-menu"
       tabindex="0"
       >
      <div>Close Menu</div>
  </div>
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
    width: 100vw;
    top: 0;
    left: 0;
    z-index: 100;
    background-color: var(--background-color);
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
  .login-box {
    display: none;
    position: fixed;
    top: calc(var(--header-height) + 10px);
    right: 10px;
    z-index: 130;
  }
  .login-box.shown {
    display: flex;
  }
  .wrapper {
    margin-top: var(--header-height);
    width: 100vw;
    overflow-x: hidden;
  }
  .overlay {
    position: fixed;
    top: var(--header-height);
    left: 0;
    z-index: 100;
    width: 100vw;
    height: 100vh;
    display: none;
    background-color: rgba(0,0,0,0.5);
  }
  .overlay:focus {
    outline: 3px dashed var(--secondary-color);
    outline-offset: -4px;
  }
  .overlay.shown {
    display: block;
  }
  .overlay div {
    display: none;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100%;
    color: var(--secondary-color);
    font-size: 2em;
    text-align: center;
    text-transform: uppercase;
    margin-left: var(--side-width);
  }
  .overlay:focus div {
    display: flex;
  }
  .main-wrapper {
    display: flex;
    justify-content: center;
    width: 100%;
  }
  main {
    max-width: var(--max-main-width);
    padding-left: 10px;
    padding-right: 10px;
  }
  aside {
    margin-top: var(--header-height);
    width: var(--side-width);
    border-right: 1px solid var(--primary-color-dimmed);
    position: fixed;
    top: 0;
    left: calc(var(--side-width) * -1);
    min-height: calc(100vh - var(--header-height));
    transition: left 0.2s ease-in-out;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    background-color: var(--background-color);
    z-index: 120;
    visibility: hidden;
  }
  aside.shown {
    left: 0;
    visibility: visible;
  }
  aside nav {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
  }
  aside nav ul {
    list-style-type: none;
    display: flex;
    margin: 0;
    padding: 15px;
    flex-wrap: wrap;
    justify-content: center;
    gap: 10px;
  }
  aside footer {
    padding: 10px;
    text-align: center;
    color: var(--text-color-dimmed);
  }
</style>
