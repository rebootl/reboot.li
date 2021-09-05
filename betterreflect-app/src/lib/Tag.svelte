<script>
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let selected = false;

</script>

<div class="tag" class:selected
     on:click={() => dispatch('click')}>
  <small><slot></slot></small>
</div>

<style>
  .tag {
    box-sizing: border-box;
    padding: 5px 5px 5px 10px;
    border: 1px solid var(--side-line-color);
    border-radius: 10px;
    min-width: 120px;
  }
  .selected {
    background-color: var(--side-selected-color);
  }
</style>
