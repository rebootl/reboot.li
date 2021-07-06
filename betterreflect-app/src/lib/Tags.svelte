<script>
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let tags = [];

  let selectedTags = new Set();

  function selectTag(tag) {
    if (selectedTags.has(tag)) {
      selectedTags.clear();
      selectedTags = selectedTags;
    } else {
      selectedTags.clear();
      selectedTags = selectedTags.add(tag);
    }
    dispatch('change', selectedTags);
  }

</script>

<div class="tags">
  {#each tags as tag}
    <div class="tag" class:selected={selectedTags.has(tag)}
         on:click={selectTag(tag)}>
      {tag}
    </div>
  {/each}
</div>

<style>
  .tags {
    display: flex;
    flex-wrap: wrap;
    max-width: 100vw;
    margin: 10px;
  }
  .tag {
    box-sizing: border-box;
    padding: 5px;
    border: 1px solid grey;
    border-radius: 5px;
  }
  .selected {
    border: 1px solid red;
    background-color: red;
  }
</style>
