<script>
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let name = '';
  export let excludes = [];

  let newItem = '';
  let newItems = [];

  function dispatchChange() {
    dispatch('change', newItems);
  }

  function addItem() {
    if (newItem === '') return;
    if (newItems.includes(newItem)) return;
    if (excludes.includes(newItem)) return;
    newItems.push(newItem);
    newItems = newItems;
    dispatchChange();
    newItem = '';
  }

  function removeItem(t) {
    newItems = newItems.filter(e => e !== t);
    dispatchChange();
  }

</script>

<div class="items">
  {#each newItems as item}
    <div class="newitem">
      <small>{item}</small>
    </div>
    <button on:click={() => removeItem(item)}>Remove</button>
  {/each}
  <input bind:value={newItem} placeholder={`New ${name}...`}>
  <button on:click={() => addItem()}>Add</button>
</div>

<style>
  .items {
    display: flex;
    flex-wrap: wrap;
    gap: 5px;
  }
  .newitem {
    background-color: var(--side-selected-color);
    border: 1px solid var(--side-line-color);
    padding: 0 5px 0 5px;
  }
</style>
