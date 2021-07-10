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
    flex-wrap: wrap;
    max-width: 100vw;
  }
  .topic {
    box-sizing: border-box;
    padding: 5px;
    /*height: 35px;*/
    border: 2px solid grey;
  }
  .selected {
    border: 2px solid red;
  }
</style>
