<script>

  export let entry = {};

</script>


{#if entry.type === 'task'}
  {entry.text}
{:else if entry.type === 'article'}
  {entry.text}
{:else if entry.type === 'link'}
  <small><a href="{entry.text}">{entry.text}</a></small><br>
  {entry.title}<br>
  {entry.comment}<br>
{:else if entry.type === 'image'}
  {#each entry.images as image}
    <img src={image.filepath} />
    <div class="imagecomment"><small>{image.comment}</small></div>
  {/each}
{:else}
  <p>oops entry type unknown: {entry.type}</p>
{/if}

<div class="tagbox">
  {#each entry.topics as topic} {topic} |{/each}<br>
  {#each entry.tags as tag} {tag} |{/each}
</div>
<hr>
