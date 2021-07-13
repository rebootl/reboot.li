<script>
  import IconButton from './IconButton.svelte';
</script>

<div class="home-button">
  <a href="/" class="home-link">
    <span class="material-icons before-icon">navigate_before</span>
    <IconButton>home</IconButton><small>Home</small>
  </a>
</div>

<style>
  .home-link {
    display: flex;
    align-items: center;
    display: flex;
    text-decoration: none;
    /*color: var(--text-color);*/
  }
  .before-icon {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 15px;
    padding-left: 10px;
  }
  small {
    text-decoration: underline;
  }
</style>
