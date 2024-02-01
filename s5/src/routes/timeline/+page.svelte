<script>

  let { data } = $props();
  // console.log(data);

</script>

<h1>Timeline</h1>

{#if data.clientData.loggedIn}
  <a href="/editPost/new">New Post...</a>
{/if}

<div class="list">
  {#each data.timelineEntries as t}
    {#if t.type === 'year'}
      <h2 class="year">{t.year}</h2>
    {:else if t.type === 'month'}
      <h3 class="month">{t.month}</h3>
    {:else}
      <div class="list-item">
        <div class="item-header">
          <small>{ t.date }</small>
          {#if t.entry?.private}
            <small><span class="material-icons">lock</span> Private</small>
          {/if}
          {#if data.clientData.loggedIn}
            <small><a href={ `/editPost/${t.entry?.id}` }><span class="material-icons">edit</span></a></small>
          {:else}
            <span></span>
          {/if}
        </div>
        <div class="item-content">
          {t.entry?.content}
          {#if t.entry?.comment}
            <small>{t.entry.comment}</small>
          {/if}
          {#if t.entry?.images}
            <div class="image-preview-box">
              {#each t.entry.images as image}
                <!--<img src={ image.path } alt={ image.comment } />-->
                <a href={ image.path }>
                  <img class="image-preview" alt={ image.comment } src={ 'data:image/png;base64,' + image.preview_data } />
                </a>
              {/each}
            </div>
          {/if}
        </div>
      </div>
    {/if}
  {:else}
    <p>No entries yet.</p>
  {/each}
</div>

<style>
  form {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }
  .list {
    margin-top: 30px;
    margin-bottom: 30px;
    display: flex;
    flex-direction: column;
    gap: 15px;
  }
  .list-item {
    display: flex;
    flex-direction: column;
    border-radius: 15px;
  }
  .list-item .item-header {
    display: flex;
    justify-content: space-between;
    color: var(--text-color-dimmed);
    font-size: 0.85em;
  }
  .list-item .item-content {
    display: flex;
    flex-direction: column;
    gap: 5px;
    margin-left: 20px;
    padding: 20px 40px 20px 40px;
    border-left: 1px solid var(--primary-color-dimmed);
  }
  small .material-icons {
    font-size: 0.85em;
  }
  .year {
    margin-top: 0.5em;
    margin-bottom: 0;
    color: var(--text-color-dimmed);
  }
  .month {
    color: var(--text-color-dimmed);
  }
  .image-preview-box {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
  }
  .image-preview {
    max-width: 120px;
    max-height: 120px;
    object-fit: contain;
  }
</style>
