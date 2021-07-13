<script>
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let topics = [];

  let selectedTopics = new Set();

  function selectTopic(topic) {
    if (selectedTopics.has(topic)) {
      selectedTopics.clear();
      selectedTopics = selectedTopics;
    } else {
      selectedTopics.clear();
      selectedTopics = selectedTopics.add(topic);
    }
    dispatch('change', selectedTopics);
  }

</script>

<div class="topics">
  {#each topics as topic}
    <div class="topic" class:selected={selectedTopics.has(topic)}
         on:click={selectTopic(topic)}>
      {topic}
    </div>
  {/each}
</div>

<style>
  .topics {
    display: flex;
    flex-flow: column;
    max-width: 100vw;
    border-top: 1px solid var(--side-line-color);
  }
  .topic {
    display: flex;
    align-items: center;
    padding: 5px 5px 5px 10px;
    height: 35px;
    border-bottom: 1px solid var(--side-line-color);
  }
  .selected {
    /*border: 2px solid red;*/
    background-color: var(--side-selected-color);
  }
</style>
