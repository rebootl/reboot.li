<script>
  import { onMount } from "svelte";
  import Entry from './Entry.svelte';

  export let entries = [];

  let limit = 5;
  let limitedEntries = entries.slice(0, limit);

  $: update(entries);

  function update() {
    limitedEntries = entries.slice(0, limit);
  }

  function addItems(e) {
    //console.log('addItems')
    if (e[e.length - 1].intersectionRatio <= 0) return;
		limitedEntries = entries.slice(0, limitedEntries.length + limit);
  }

  function initInfiniteScroll() {
    const bottomObserver = new IntersectionObserver(
      (e) => { addItems(e) },
      { threshold: 0.5 }
    );

    const liMutationObserver = new MutationObserver(
      (m, o) => updateObserver(m, o)
    );

    function updateObserver(m, o) {
      const triggerElement = document.querySelector('.triggerelement');
      if (triggerElement) {
        triggerElement.classList.remove('triggerelement');
        bottomObserver.unobserve(triggerElement);
      }
      const el = document.querySelector('.entrieslist');
      if (!el) return;
      const newTriggerElement = el.children[el.children.length - 2];
      if (newTriggerElement) {
        //console.log('add observer')
        newTriggerElement.classList.add('triggerelement')
        bottomObserver.observe(newTriggerElement);
      }
    }

    const li = document.querySelector('.entrieslist');
    liMutationObserver.observe(li, { childList: true });
    updateObserver([], [])
  }

  onMount(async () => {
    initInfiniteScroll();
	});
</script>

<div class="entrieslist">
  {#each limitedEntries as entry}
    <div class="entry">
      <Entry {entry} />
    </div>
  {/each}
</div>

<style>
</style>
