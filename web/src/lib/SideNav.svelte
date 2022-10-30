<script>
  import { createEventDispatcher } from 'svelte';
  import { currentTopics, currentTags, currentTagsByTopics, showMenu }
    from '$lib/store';
  import { refs } from '$lib/refs';

  import Topic from '$lib/Topic.svelte';
  import Tag from '$lib/Tag.svelte';
  import BackButton from '$lib/BackButton.svelte';
  import HeaderLinks from '$lib/HeaderLinks.svelte';

  const dispatch = createEventDispatcher();

  export let entries = [];
  export let showOnWide = true;
  export let backbutton = false;
  export let ref = 'note';
  console.log(ref)

  let topics = [];
  let tags = [];
  let selectedTopics = [];
  let selectedTags = [];
  let tagsByTopics = {};
  let filterText = '';

  $: setTopicsTags(entries);
  $: dispatchChange(filterText);

  function dispatchChange() {
    dispatch('change', [ selectedTopics, selectedTags, filterText ]);
  }

  function setTopicsTags() {
    entries.forEach(entry => {
      entry.topics.forEach(topic => {
        if (!topics.includes(topic)) topics.push(topic)
        if (!tagsByTopics.hasOwnProperty(topic))
          // must create copy here!!
          tagsByTopics[topic] = [ ...entry.tags ];
        else
          entry.tags.forEach(tag => {
            if (!tagsByTopics[topic].includes(tag))
              tagsByTopics[topic].push(tag);
        });
      });
    });
    topics = topics.sort();
    if (!backbutton) {
      $currentTagsByTopics = tagsByTopics;
      $currentTopics = topics;
    }
    setTags();
  }

  function setTags() {
    selectedTags = [];
    const r = [];
    if (selectedTopics.length > 0) {
      selectedTopics.forEach(topic => {
        tagsByTopics[topic].forEach(tag => {
          if (!r.includes(tag)) r.push(tag);
        });
      });
    } else {
      entries.forEach(e => {
        e.tags.forEach(tag => {
          if (!r.includes(tag)) r.push(tag);
        });
      });
    }
    tags = r.sort();
    if (!backbutton) {
      $currentTags = tags;
    }
  }

  function selectTopic(topic) {
    if (selectedTopics.includes(topic)) {
      selectedTopics = [];
    } else {
      selectedTopics = [ topic ];
    }
    setTags();
    if ($showMenu) $showMenu = false;
    dispatch('change', [ selectedTopics, selectedTags, filterText ]);
  }

  function selectTag(tag) {
    if (selectedTags.includes(tag)) {
      selectedTags = [];
    } else {
      selectedTags = [ tag ];
    }
    if ($showMenu) $showMenu = false;
    dispatch('change', [ selectedTopics, selectedTags, filterText ]);
  }

  function click() {
    if ($showMenu) $showMenu = false;
  }
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div class="overlay" class:show={$showMenu}
     on:click={() => $showMenu = !$showMenu}></div>
<aside class:show={$showMenu} class:hidden={!showOnWide}>
  {#if backbutton}
    <BackButton href={refs[ref].href} icon={refs[ref].icon}
                on:click={() => click()}>{refs[ref].text}</BackButton>
  {/if}
  <div class="header-links">
    <HeaderLinks side={true} />
  </div>
  <div class="padding">
    {#if !backbutton}
      <div class="quickfilter">
        <input bind:value={filterText} placeholder="Quickfilter entries..."
               class="quickfilter-input">
      </div>
    {/if}
    <div class="items">
      {#each topics as topic}
        <Topic selected={selectedTopics.includes(topic)}
               on:click={() => selectTopic(topic)}>
          {topic}
        </Topic>
      {/each}
    </div>
    <div class="items">
      {#each tags as tag}
        <Tag selected={selectedTags.includes(tag)}
               on:click={() => selectTag(tag)}>
          {tag}
        </Tag>
      {/each}
    </div>
  </div>
</aside>

<style>
  aside {
    width: var(--side-width);
    position: absolute;
    top: var(--header-height);
    left: 0;
    background-color: var(--background-color);
  }
  aside.hidden {
		display: none;
	}
	@media all and (max-width: 600px) { /* 1000px = side width + max. main width */
		aside.hidden {
			display: initial;
		}
	}
  .padding {
    padding: 35px 15px 15px 15px;
    display: flex;
    flex-flow: column;
    gap: 35px;
  }
  .items {
    display: flex;
    flex-flow: column;
    flex-wrap: wrap;
    align-items: flex-start;
    gap: 10px;
  }
  .quickfilter-input {
    width: 160px;
  }
  .overlay {
    display: none;
    position: fixed;
    left: 0;
    top: 0;
    width: 100vw;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
  }
  .header-links {
    display: none;
    padding: 15px;
    border-bottom: 1px solid var(--side-line-color);
  }
  @media all and (max-width: 600px) { /* 1000px = side width + max. main width */
    aside {
      position: fixed;
      left: calc(-1 * 80vw);
      overflow: scroll;
      height: calc(100vh - var(--header-height));
      width: 80vw;
    }
    .show {
      left: 0;
    }
    .overlay.show {
      display: initial;
    }
    .header-links {
      display: flex;
      justify-content: center;
    }
    .padding {
      padding-top: 15px;
      gap: 15px;
    }
  }
  /*.hidden {
    left: -var(--side-width);
  }*/
</style>
