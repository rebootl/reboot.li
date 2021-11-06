<script>
  import IconButton from './IconButton.svelte';

  export let href = '';
  export let icon = '';

</script>

<div class="back-button">
  <a {href} class="back-link">
    <span class="material-icons before-icon">navigate_before</span>
    <IconButton>{icon}</IconButton><small><slot></slot></small>
  </a>
</div>

<style>
  .back-link {
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
    /*text-decoration: underline;*/
  }
</style>
