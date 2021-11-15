<script>
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();
</script>

<div class="icon-button" on:click={dispatch('click')}>
  <span class="material-icons"><slot></slot></span>
</div>

<style>
  .icon-button {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 40px;
    width: 40px;
  }
</style>
