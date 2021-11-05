<svelte:head>
	<title>reboot.li - a personal website</title>
</svelte:head>

<script>
  import { page } from '$app/stores';
  import { onMount } from 'svelte';

  import Login from '$lib/Login.svelte';

  let showLogin = false;

</script>

<div class="wrapper">
  <header>
    <!--<div class="icon-button">
      <a href="/" title="Home">
        <span class="material-icons">home</span>
      </a>
    </div>-->
    <!--<div class="icon-button"></div>-->
    <div class="logo-box">
      <img class="logo" alt="Logo" src="/static/betterreflect-app-logo.png" />
    </div>

    <div class="right-box">
      <div class="header-links">
        <a href="/" class:active={'/' === $page.path}>Home</a>
        <a href="/tasks" class:active={'/tasks' === $page.path}>Tasks</a>
        <a href="/notes" class:active={'/notes' === $page.path}>Notes</a>
        <a href="/links" class:active={'/links' === $page.path}>Links</a>
        <a href="/images" class:active={'/images' === $page.path}>Images</a>
				<a href="/about" class:active={'/about' === $page.path}>About</a>
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
	:global(body) {

    --primary-color: #5af0a4;
    --secondary-color: #c9a6a4;

    --background-color: #1f2126;
    --text-color: #fff;

    --header-background-color: #141417;

    --dialog-background-color: #383b46;

    --side-selected-color: #383b46;
    --side-line-color: #4f5260;

    --main-line-color: #4f5260;
    --main-topic-label-background-color: #4f5260;
    --main-topic-label-text-color: #fff;
    --main-text-label-color: #adadbc;

    --error-color: #ff0000;
    --on-error-color: #fff;

    --header-height: 55px;
    --side-width: 200px;
		--max-main-width: 600px;

    margin: 0;
    background-color: var(--background-color);
    color: var(--text-color);
    font-family: sans-serif;
    line-height: 1.5em;
  }
  :global(a) {
    color: var(--primary-color);
  }
  :global(h2) {
    color: var(--text-color);
  }
  :global(@font-face) {
    font-family: 'Material Icons';
    font-style: normal;
    font-weight: 400;
    src: url(/static/iconfont/MaterialIcons-Regular.eot); /* For IE6-8 */
    src: local('/static/iconfont/Material Icons'),
      local('/static/iconfont/MaterialIcons-Regular'),
      url(/static/iconfont/MaterialIcons-Regular.woff2) format('woff2'),
      url(/static/iconfont/MaterialIcons-Regular.woff) format('woff'),
      url(/static/iconfont/MaterialIcons-Regular.ttf) format('truetype');
  }
  :global(.material-icons) {
    font-family: 'Material Icons';
    font-weight: normal;
    font-style: normal;
    font-size: 24px;  /* Preferred icon size */
    display: inline-block;
    line-height: 1;
    text-transform: none;
    letter-spacing: normal;
    word-wrap: normal;
    white-space: nowrap;
    direction: ltr;

    /* Support for all WebKit browsers. */
    -webkit-font-smoothing: antialiased;
    /* Support for Safari and Chrome. */
    text-rendering: optimizeLegibility;

    /* Support for Firefox. */
    -moz-osx-font-smoothing: grayscale;

    /* Support for IE. */
    font-feature-settings: 'liga';
  }
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
  .header-links {
    display: flex;
    align-items: center;
    gap: 10px;
  }
  .header-links a {
    text-decoration: none;
  }
  .header-links a.active {
    text-decoration: underline;
  }
  .icon-button {
    display: flex;
    justify-content: center;
    align-items: center;
    width: var(--header-height);
    height: var(--header-height);
  }
  .icon-button a {
    color: var(---text-color);
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
</style>
