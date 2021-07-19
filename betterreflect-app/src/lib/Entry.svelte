<script>
  import moment from 'moment';

  //const md = window.markdownit();
  const dateFormat = 'MMM D YYYY - HH:mm';

  export let entry = {};

  let date = '';

  $: update(entry);

  function update() {
    date = moment(new Date(entry.date)).format(dateFormat);
  }

</script>

<div class="entry">
  <div class="entry-header"><small>{date}</small></div>

  <div class="entry-content">
    {#if entry.type === 'task'}
      {entry.text}
    {:else if entry.type === 'article'}
      {entry.text}
    {:else if entry.type === 'link'}
      <small><a href="{entry.text}">{entry.text}</a></small><br>
      {entry.title}<br>
      {entry.comment}
    {:else if entry.type === 'image'}
      {#each entry.images as image}
        <img src={image.filepath} />
        <div class="imagecomment"><small>{image.comment}</small></div>
      {/each}
    {:else}
      <p>oops entry type unknown: {entry.type}</p>
    {/if}
  </div>

  <div class="tagbox">
    {#each entry.topics as topic}
      <small class="topic-label">{topic}</small>
    {/each}
    {#each entry.tags as tag}
      <small class="tag-label">{tag}</small>
    {/each}
  </div>
</div>

<style>
  .entry {
    border-bottom: 1px solid var(--main-line-color);
    padding: 10px 0 10px 0;
  }
  .entry-header {
    color: var(--main-text-label-color);
    font-size: 0.85em;
  }
  .entry-content {
    margin: 10px 0 20px 0;
  }
  .tagbox {
    display: flex;
    gap: 5px;
  }
  .topic-label {
    padding: 0 5px 0 5px;
    background-color: var(--main-topic-label-background-color);
    color: var(--main-topic-label-text-color);
  }
  .tag-label {
    padding: 0 5px 0 5px;
    border: 1px solid var(--main-line-color);
    border-radius: 10px;
    color: var(--main-text-label-color);
  }
</style>
