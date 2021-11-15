<script>
  import Tag from '$lib/Tag.svelte';
  import NewItems from '$lib/NewItems.svelte';
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let tagsByTopics = [];
  export let newTopics = [];
  export let selectedItems = [];
  export let loadTopics = [];

  let items = [];
  let newItems = [];

  $: updateTags(newTopics);
  $: loadEdit(loadTopics);

  // -> use onmount instead?
  function loadEdit() {
    if (loadTopics) {
      newTopics = loadTopics;
      updateTags();
    }
  }

  function updateTags() {
    const r = [];
    if (newTopics.length === 0) {
      Object.keys(tagsByTopics)
        .forEach(topic => tagsByTopics[topic]
        .forEach(tag => {
          if (!r.includes(tag)) r.push(tag);
        }
      ));
    } else {
      newTopics.forEach(topic => {
        if (!tagsByTopics.hasOwnProperty(topic)) return;
        tagsByTopics[topic].forEach(tag => {
          if (!r.includes(tag)) r.push(tag);
        });
      });
      //selectedItems = selectedItems.filter(e => items.includes(e));
      selectedItems.forEach(e => {
        if (!r.includes(e)) r.push(e);
      });
    }
    items = r.sort();
    dispatchChange();
  }

  function selectItem(item) {
    if (newItems.includes(item)) return;
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

<NewItems excludes={items} name="Tag"
          on:change={(e) => newItemsChanged(e.detail)} />
<div class="items">
  {#each items as item}
    <Tag selected={selectedItems.includes(item)}
         on:click={() => selectItem(item)}>
      {item}
    </Tag>
  {/each}
</div>

<style>
  .items {
    display: flex;
    flex-wrap: wrap;
    gap: 5px;
  }
</style>
