<script>
  import Login from '$lib/Login.svelte';

  let showLogin = false;

</script>

<div class="grid-box">
  <header>
    <a href="/" title="Home">
      <img class="logo" alt="Logo" src="/static/betterreflect-app-logo.png" />
    </a>

    <div class="login-button" on:click={() => showLogin = !showLogin}>
      <span class="material-icons">account_circle</span>
    </div>
    <div class="login-box" class:show={showLogin}>
      <Login />
    </div>
  </header>

  <div class="content">
    <slot></slot>
  </div>
</div>

<style>
	:global(body) {

    --primary-color: #5af0a4;

    --background-color: #3b3e48;
    --text-color: #fff;

    --header-background-color: #141417;

    --dialog-background-color: #3e646b;

    --header-height: 48px;

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
  .grid-box {
    display: grid;
  }
  header {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    display: flex;
    justify-content: space-between;
    height: var(--header-height);
    background-color: var(--header-background-color);
  }
  .logo {
    width: var(--header-height);
    height: var(--header-height);
  }
  .login-button {
    display: flex;
    justify-content: center;
    align-items: center;
    width: var(--header-height);
    height: var(--header-height);
  }
  .login-box {
    display: none;
    position: absolute;
    top: 55px;
    right: 10px;
  }
  .login-box.show {
    display: flex;
  }
  .content {
    margin-top: var(--header-height);
  }
</style>
