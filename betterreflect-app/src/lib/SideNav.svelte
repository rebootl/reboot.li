<script>
  import Topic from '$lib/Topic.svelte';
  import Tag from '$lib/Tag.svelte';
  import BackButton from '$lib/BackButton.svelte';
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let entries = [];
  export let hidden = false;
  export let backbutton = false;
  export let ref = '';

  let topics = [];
  let tags = [];
  let selectedTopics = [];
  let selectedTags = [];
  let tagsByTopics = {};

  let refs = {
    article: {
      href: '/notes',
      text: 'Notes',
      icon: 'list'
    },
    note: {
      href: '/notes',
      text: 'Notes',
      icon: 'list'
    },
    link: {
      href: '/links',
      text: 'Links',
      icon: 'list'
    },
    task: {
      href: '/tasks',
      text: 'Tasks',
      icon: 'list'
    },
    image: {
      href: '/images',
      text: 'Images',
      icon: 'list'
    },
    home: {
      href: '/',
      text: 'Home',
      icon: 'home'
    },
  }

  $: setTopicsTags(entries);

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
  }

  function selectTopic(topic) {
    if (selectedTopics.includes(topic)) {
      selectedTopics = [];
    } else {
      selectedTopics = [ topic ];
    }
    setTags();
    dispatch('change', [ selectedTopics, selectedTags ]);
  }

  function selectTag(tag) {
    if (selectedTags.includes(tag)) {
      selectedTags = [];
    } else {
      selectedTags = [ tag ];
    }
    dispatch('change', [ selectedTopics, selectedTags ]);
  }

</script>

<aside class:hidden={hidden}>
  {#if backbutton}
    <BackButton href={refs[ref].href} icon={refs[ref].icon}>{refs[ref].text}</BackButton>
  {/if}
  <div class="padding">
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
  .hidden {
    left: -var(--side-width);
  }
</style>
