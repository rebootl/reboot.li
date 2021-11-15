<script>
  import Tag from './Tag.svelte';
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
    <Tag selected={selectedTags.has(tag)}
         on:click={() => selectTag(tag)}>
      {tag}
    </Tag>
  {/each}
</div>

<style>
  .tags {
    display: flex;
    flex-flow: column;
    margin: 10px 15px 10px 15px;
    gap: 5px;
  }
</style>
