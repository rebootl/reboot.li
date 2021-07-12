<script>
  import Login from '$lib/Login.svelte';

  let showLogin = false;

</script>

<div class="grid-box">
  <header>
    <!--<div class="icon-button">
      <a href="/" title="Home">
        <span class="material-icons">home</span>
      </a>
    </div>-->
    <div class="icon-button"></div>
    <div class="logo-box">
      <img class="logo" alt="Logo" src="/static/betterreflect-app-logo-t3.png" />
    </div>

    <div class="icon-button" on:click={() => showLogin = !showLogin}>
      <span class="material-icons">account_circle</span>
    </div>
    <div class="login-box" class:show={showLogin}>
      <Login />
    </div>
  </header>

  <div class="routed-content">
    <slot></slot>
  </div>
</div>

<style>
	:global(body) {

    --primary-color: #5af0a4;

    --background-color: #3b3e48;
    --text-color: #fff;

    --header-background-color: #141417;

    --dialog-background-color: #4f5260;

    --header-height: 48px;
    --side-width: 220px;

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
    height: 100vh;
    grid-template-rows: auto 1fr; /*var(--header-height) auto;*/
    grid-template-areas:
      "header"
      "routed-content";
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
    width: calc(var(--header-height) - 10px);
    height: calc(var(--header-height) - 10px);
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
    top: 55px;
    right: 10px;
  }
  .login-box.show {
    display: flex;
  }
  .routed-content {
    grid-area: routed-content;
    overflow-y: scroll;
  }
</style>
