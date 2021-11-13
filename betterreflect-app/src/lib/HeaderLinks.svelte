<script>
  import { page } from '$app/stores';

  export let side = false;

</script>

<div class="header-links" class:column={side}>
  <a href="/" class:active={'/' === $page.path}>Home</a>
  <a href="/tasks" class:active={'/tasks' === $page.path}>Tasks</a>
  <a href="/notes" class:active={'/notes' === $page.path}>Notes</a>
  <a href="/links" class:active={'/links' === $page.path}>Links</a>
  <a href="/images" class:active={'/images' === $page.path}>Images</a>
  <a href="/about" class:active={'/about' === $page.path}>About</a>
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
