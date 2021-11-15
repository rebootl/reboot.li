<script>
  import Topic from './Topic.svelte';
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
    <Topic selected={selectedTopics.has(topic)}
           on:click={() => selectTopic(topic)}>
      {topic}
    </Topic>
  {/each}
</div>

<style>
  .topics {
    display: flex;
    flex-flow: column;
    max-width: 100vw;
    border-top: 1px solid var(--side-line-color);
  }
</style>
