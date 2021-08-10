<script>
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let topics = [];

  let selectedTopics = new Set();
  let newTopics = new Set();

  function selectTopic(topic) {
    if (selectedTopics.has(topic)) {
      selectedTopics.delete(topic);
    } else {
      selectedTopics.add(topic);
    }
    selectedTopics = selectedTopics;
    dispatch('change', selectedTopics);
  }

</script>

<div class="topics">
  {#each topics as topic}
    <div class="topic" class:selected={selectedTopics.has(topic)}
         on:click={selectTopic(topic)}>
      <small>{topic}</small>
    </div>
  {/each}
</div>

<style>
  .topics {
    display: flex;
    flex-wrap: wrap;
    gap: 5px;
    /*flex-flow: column;*/
    /*max-width: 100vw;*/
    /*border-top: 1px solid var(--side-line-color);*/
  }
  .topic {
    display: flex;
    align-items: center;
    padding: 0 5px 0 5px;
    height: 35px;
    border: 1px solid var(--side-line-color);
    border-radius: 5px;
  }
  .selected {
    /*border: 2px solid red;*/
    background-color: var(--side-selected-color);
  }
</style>
