<script>
  import { page } from '$app/stores';
  import { showMenu } from '$lib/store';

  export let side = false;

  function click() {
    if (side) $showMenu = false;
  }
</script>

<div class="header-links" class:column={side}>
  <a href="/" class:active={'/' === $page.url.pathname} on:click={() => click()}>Home</a>
  <a href="/tasks" class:active={'/tasks' === $page.url.pathname} on:click={() => click()}>Tasks</a>
  <a href="/notes" class:active={'/notes' === $page.url.pathname} on:click={() => click()}>Notes</a>
  <a href="/links" class:active={'/links' === $page.url.pathname} on:click={() => click()}>Links</a>
  <a href="/images" class:active={'/images' === $page.url.pathname} on:click={() => click()}>Images</a>
  <a href="/about" class:active={'/about' === $page.url.pathname} on:click={() => click()}>About</a>
</div>

<style>
  .header-links {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    gap: 10px;
  }
  .column {
    /*flex-flow: column;*/
    gap: 15px;
  }
  .header-links a {
    text-decoration: none;
  }
  .header-links a.active {
    text-decoration: underline;
  }
</style>
