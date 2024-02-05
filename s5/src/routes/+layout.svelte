<svelte:head>
	<title>reboot.li - my personal website</title>
</svelte:head>

<script>
  import 'material-icons/iconfont/material-icons.css';
	import '../global.css';
  import Login from '$lib/Login.svelte';

  /** @typedef clientData
    * @property {boolean} loggedIn
    * @property {string | null} username
    */
  /**
   * @type {{clientData: clientData}}
   */
  export let data;
  // console.log('data layout', data);

  /**
   * @type {boolean}
   */
  let showMenu = false;
  /**
   * @type {boolean}
   */
  let showLogin = false;  
</script>

<header>
  <button class="icon-button" onclick={() => showMenu = !showMenu}>
		<span class="material-icons">menu</span>
  </button>
  <div class="logo-box">
    <img class="logo" alt="Logo" src="/logo.png" />
  </div>
  <button class="icon-button" onclick={() => showLogin = !showLogin}>
    <span class="material-icons">account_circle</span>
  </button>
</header>
<div class="wrapper">
  <aside class:shown={showMenu}>
    <nav>
      <a href="/" onclick={() => showMenu = false}>Home</a> | 
      <a href="/timeline" onclick={() => showMenu = false}>Timeline</a> | 
      <a href="/notes" onclick={() => showMenu = false}>Notes</a> |
      <a href="/links" onclick={() => showMenu = false}>Links</a>
    </nav>
    <div class="credits">
      <small>Copyright 2024 Cem Aydin<br />
        Created with <a href="https://kit.svelte.dev/" target="_blank">SvelteKit</a></small>
    </div>
  </aside>
  <div class="overlay" class:shown={showMenu || showLogin}
       onclick={() => {showMenu = false ; showLogin = false}}
       onkeydown={()=>{}}
       role="dialog"
       aria-disabled="true"></div>
  <div class="main-wrapper" class:menu-shown={showMenu}>
    <main>
      <slot></slot>
    </main>
    <footer>
    </footer>
  </div>
  <div class="login-box" class:shown={showLogin}>
    <Login clientData={data.clientData} />
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
  .overlay.shown {
    display: block;
  }
  .main-wrapper {
    display: flex;
    justify-content: center;
    width: 100%;
    /*margin-left: 0;
    transition: margin-left 0.2s ease-in-out, width 0.2s ease-in-out;*/
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
  }
  aside.shown {
    left: 0;
  }
  aside nav {
    display: flex;
    gap: 10px;
    padding: 20px;
    flex-wrap: wrap;
  }
  div.credits {
    padding: 10px;
    text-align: center;
    color: var(--text-color-dimmed);
  }
/*  @media (max-width: 660px) {
    .main-wrapper {
      min-width: 300px;
    }
    main {
      min-width: 300px;
    }
  }*/
  /*
  @media (max-width: 400px) {
    .main-wrapper {
      min-width: 200px;
    }
    main {
      min-width: 200px;
    }
  }
  */
</style>
