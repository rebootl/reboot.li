<script>
  import Topic from '$lib/Topic.svelte';
  import NewItems from '$lib/NewItems.svelte';
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let items = [];
  export let selectedItems = [];

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

  function newItemsChanged(v) {
    newItems = v;
    dispatchChange();
  }

</script>

<NewItems excludes={items} name="Topic"
          on:change={(e) => newItemsChanged(e.detail)} />
<div class="items">
  {#each items as item}
    <Topic type="edit" selected={selectedItems.includes(item)}
           on:click={selectItem(item)}>
      {item}
    </Topic>
  {/each}
</div>

<style>
  .items {
    display: flex;
    flex-wrap: wrap;
    gap: 5px;
  }
</style>
