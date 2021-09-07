<script>
  import EditTypes from './EditTypes.svelte';
  import EditTopics from './EditTopics.svelte';
  import EditTags from './EditTags.svelte';

  export let topics = [];
  export let tagsByTopics = {};

  let showAddElements = true;

  let text = '';
  let newTopics = [];
  let tags = [];
  let newTags = [];

  $: textInput(text)

  function textInput() {
    if (text === '') showAddElements = true;
    else showAddElements = true;
    console.log(text)
  }

  function setNewTopics(v) {
    newTopics = v;
    console.log(newTopics)
  }

  function setNewTags(v) {
    newTags = v;
    console.log(newTags)
  }

</script>

<div class="newentry-box">

  <div>
    <textarea class="newentry-text"
              placeholder="New Entry..."
              bind:value={text}></textarea>
  </div>
  {#if showAddElements}
    <EditTypes />
    <EditTopics items={topics} name="Topic"
                on:change={(e) => setNewTopics(e.detail)} />
    <EditTags {tagsByTopics} {newTopics} on:change={(e) => setNewTags(e.detail)} />

  {/if}

</div>

<style>
  .newentry-box {
    display: flex;
    flex-flow: column;
    gap: 20px;
    padding: 20px 0 20px 0;
    border-bottom: 1px solid var(--main-line-color);
  }
  .newentry-text {
    width: 170px;
    height: 20px;
    padding: 10px;
  }
</style>
