<script>
  import moment from 'moment';
  import marked from 'marked';
  import { session } from '$app/stores';

  const dateFormat = 'MMM D YYYY - HH:mm';

  export let entry = {};

  let date = '';
  let html = '';
  let url = '';

  $: update(entry);

  function update() {
    date = moment(new Date(entry.date)).format(dateFormat);

    url = '/entry/' + entry.id;

    if (entry.type === 'task' || entry.type === 'article')
      html = marked(entry.text);
  }

</script>

<div class="entry">
  <div class="entry-header">
    <small>
      {date}
      <a href={url}>
        <span class="material-icons header-icon">link</span>
      </a>
      {#if entry.pinned}
        <span class="material-icons header-icon">adjust</span>
        <!-- push_pin icon not working for some reason
        -->
        Pinned
      {/if}
      {#if entry.private}
        <span class="material-icons header-icon">lock</span>
        Private
      {/if}
    </small>
    <small>
      {#if $session.loggedIn}
        <a href={url + '?edit'}>
          <span class="material-icons header-icon">edit</span>
        </a>
      {/if}
    </small>
  </div>

  <div class="entry-content">
    {#if entry.type === 'task'}
      {@html html}
    {:else if entry.type === 'article'}
      {@html html}
    {:else if entry.type === 'link'}
      <small><a href="{entry.text}">{entry.text}</a></small><br>
      {entry.title}<br>
      {entry.comment}
    {:else if entry.type === 'image'}
      {#if entry.images}
        {#each entry.images as image}
          <img src={image.filepath} />
          <div class="imagecomment"><small>{image.comment}</small></div>
        {/each}
      {/if}
    {:else}
      <p>oops entry type unknown: {entry.type}</p>
    {/if}
  </div>

  <div class="tagbox">
    {#each entry.topics as topic}
      <small class="topic-label label">{topic}</small>
    {/each}
    {#each entry.tags as tag}
      <small class="tag-label label">{tag}</small>
    {/each}
  </div>
</div>

<style>
  .entry {
    border-bottom: 1px solid var(--main-line-color);
    padding: 10px 0 10px 0;
    min-height: 120px; /* set for inf. scroll loading calc. */
  }
  .entry-header {
    display: flex;
    justify-content: space-between;
    color: var(--main-text-label-color);
    font-size: 0.85em;
  }
  .header-icon {
    font-size: 1.2em;
    margin-left: 5px;
    vertical-align: sub;
  }
  .entry-content {
    margin: 10px 0 20px 0;
  }
  img {
    max-width: 100%;
    max-height: 90vh;
  }
  .tagbox {
    display: flex;
    gap: 5px;
  }
  .label {
    padding: 0 5px 0 5px;
  }
  .topic-label {
    background-color: var(--main-topic-label-background-color);
    color: var(--main-topic-label-text-color);
  }
  .tag-label {
    border: 1px solid var(--main-line-color);
    border-radius: 10px;
    color: var(--main-text-label-color);
  }
</style>
