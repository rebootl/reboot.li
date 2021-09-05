<script>
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let selected = false;

</script>

<div class="topic" class:selected
     on:click={() => dispatch('click')}>
  <slot></slot>
</div>

<style>
  .topic {
    display: flex;
    align-items: center;
    padding: 5px 10px 5px 10px;
    height: 35px;
    border-bottom: 1px solid var(--side-line-color);
    min-width: 120px;
  }
  .selected {
    background-color: var(--side-selected-color);
  }
</style>
