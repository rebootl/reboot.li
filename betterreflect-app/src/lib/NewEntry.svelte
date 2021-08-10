<script>
  import NewTopics from './NewTopics.svelte';

  export let topics = [];
  export let tags = [];

  let showAddElements = true;

  let text = '';

  $: textInput(text)

  function textInput() {
    if (text === '') showAddElements = true;
    else showAddElements = true;
    console.log(text)
  }

</script>

<div class="newentry-box">

  <div>
    <textarea class="newentry-text"
              placeholder="New Entry..."
              bind:value={text}></textarea>
  </div>
  {#if showAddElements}
    <div>
      Topics/Tags:
      <NewTopics {topics} />
    </div>
  {/if}

</div>

<style>
  .newentry-box {
    margin: 20px 0 20px 0;
    border-bottom: 1px solid var(--main-line-color);
  }
  .newentry-text {
    width: 170px;
    height: 20px;
    padding: 15px;
    background-color: var(--dialog-background-color);
    color: var(--text-color);
    border: 0;
  }
</style>
