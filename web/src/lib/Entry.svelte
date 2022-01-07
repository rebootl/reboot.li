<script>
  import moment from 'moment';
  import { marked } from 'marked';
  import hljs from 'highlight.js';
  import 'highlight.js/styles/atom-one-dark.css';

  marked.setOptions({
    highlight: function(code, lang) {
      const language = hljs.getLanguage(lang) ? lang : 'plaintext';
      return hljs.highlight(code, { language: language }).value;
    },
    langPrefix: 'hljs language-', // highlight.js css expects a top-level 'hljs' class.
  });

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

    if ([ 'task', 'note', 'news'].includes(entry.type)) {
      html = marked.parse(entry.text);
    }
  }
</script>

<div class="entry" class:news={entry.type === 'news'}>
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
    {#if [ 'task', 'note', 'news'].includes(entry.type)}
      {@html html}
    {:else if entry.type === 'link'}
      <small><a href="{entry.text}">{entry.text}</a></small><br>
      {entry.title}<br>
      <small>{entry.comment}</small>
    {:else if entry.type === 'image'}
      {#if entry.images}
        {#each entry.images as image}
          <img src={image.url} alt="entry" />
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
    padding: 10px 0 20px 0;
    min-height: 120px; /* set for inf. scroll loading calc. */
  }
  .news {
    padding-bottom: 5px;
    min-height: 60px; /* set for inf. scroll loading calc. */
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
  .tagbox {
    display: flex;
    flex-wrap: wrap;
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
