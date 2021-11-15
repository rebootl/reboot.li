<script>
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let selected = false;
  export let type = '';

</script>

<div class="topic" class:selected class:edit={type === 'edit'}
     on:click={() => dispatch('click')}>
  <slot></slot>
</div>

<style>
  .topic {
    display: flex;
    align-items: center;
    padding: 5px 10px 5px 10px;
    height: 35px;
    border: 1px solid var(--side-line-color);
  }
  .edit {
    border: 1px solid var(--side-line-color);
    width: 120px;
  }
  .selected {
    background-color: var(--side-selected-color);
  }
</style>
