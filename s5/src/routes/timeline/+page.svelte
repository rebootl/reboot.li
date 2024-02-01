<script>
  import dayjs from 'dayjs';

  import { onMount } from 'svelte';

	// export let data;
  let { data } = $props();
  // console.log(data);
  
	let entries = data.entries;
  // console.log(entries);

  let previousYear = dayjs(entries[0].created_at).format('YYYY');
  let previousMonth = dayjs(entries[0].created_at).format('MMMM');
  
  // we want to display a timeline with entries grouped by year and month
  // we do this by creating a new array with entries and year/month headers
  // we also sort the entries by date
  
  /** @typedef {Object} TimelineEntry
    * @property {'year' | 'month' | 'entry'} type
    * @property {string | null} year
    * @property {string | null} month
    * @property {import('$lib/server/db.js').EntryData | null} entry
    * @property {string | null} date
    */
  /** @type {TimelineEntry[]} */
  let timelineEntries = $state([
    {
      type: 'year',
      year: previousYear,
      month: null,
      entry: null,
      date: null,
    },
    {
      type: 'month',
      year: null,
      month: previousMonth,
      entry: null,
      date: null,
    },
  ]);

  onMount(() => {
    const sortedEntries = entries.sort((a, b) => {
      const aDate = a.manual_date || a.created_at;
      const bDate = b.manual_date || b.created_at;
      return dayjs(aDate).isBefore(dayjs(bDate)) ? 1 : -1;
    });

    for (const entry of sortedEntries) {
      const entryDate = entry.manual_date || entry.created_at;
  
      const entryYear = dayjs(entryDate).format('YYYY');
      const entryMonth = dayjs(entryDate).format('MMMM');
      // console.log(entryYear);
      // console.log(entryMonth);
      if (entryYear !== previousYear) {
        timelineEntries.push({
          type: 'year',
          year: entryYear,
          month: null,
          entry: null,
          date: null,
        });
        previousYear = entryYear;
      }
      if (entryMonth !== previousMonth) {
        timelineEntries.push({
          type: 'month',
          year: null,
          month: entryMonth,
          entry: null,
          date: null,
        });
        previousMonth = entryMonth;
      }
      timelineEntries.push({
        type: 'entry',
        year: null,
        month: null,
        entry: entry,
        date: entryDate,
      });
    }
    // console.log(timelineEntries);
  });

</script>

<h1>Timeline</h1>

{#if data.clientData.loggedIn}
  <a href="/editPost/new">New Post...</a>
{/if}

<div class="list">
  {#each timelineEntries as t}
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
