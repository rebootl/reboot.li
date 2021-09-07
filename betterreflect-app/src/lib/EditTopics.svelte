<script>
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let name = 'Items';
  export let items = [];

  let selectedItems = [];
  let newItem = '';
  let newItems = [];

  function selectItem(item) {
    if (selectedItems.includes(item)) {
      selectedItems = selectedItems.filter(e => e !== item);
    } else {
      selectedItems.push(item);
      selectedItems = selectedItems;
    }
    dispatchChange();
  }

  function dispatchChange() {
    dispatch('change', [ ...selectedItems, ...newItems ]);
  }

  function addItem() {
    if (newItem === '') return;
    if (newItems.includes(newItem)) return;
    if (items.includes(newItem)) return;
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
<div class="items">
  {#each items as item}
    <div class="item" class:selected={selectedItems.includes(item)}
         on:click={selectItem(item)}>
      {item}
    </div>
  {/each}
</div>

<style>
  .items {
    display: flex;
    flex-wrap: wrap;
    gap: 3px;
  }
  .item {
    display: flex;
    align-items: center;
    padding: 5px 10px 5px 10px;
    height: 35px;
    width: 120px;
    border: 1px solid var(--side-line-color);
  }
  .newitem {
    background-color: var(--side-selected-color);
    border: 1px solid var(--side-line-color);
    padding: 0 5px 0 5px;
  }
  .selected {
    background-color: var(--side-selected-color);
  }
</style>
