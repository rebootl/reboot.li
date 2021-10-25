<script>
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let selectedType = 'any';

  let types = [
    {
      icon: 'view_module',
      label: 'Any',
			type: 'any'
    },
    {
      icon: 'task_alt',
      label: 'Tasks',
			type: 'task'
    },
    {
      icon: 'link',
      label: 'Links',
			type: 'link'
    },
    {
      icon: 'article',
      label: 'Articles',
			type: 'article'
    },
    {
      icon: 'image',
      label: 'Images',
			type: 'image'
    }
  ];

  function selectType(t) {
    selectedType = t;
    dispatch('change', t);
  }

//           checked={type.type === 'all'}
</script>

<div class="types">
  Types:
  {#each types as type}
    <input type="radio" name="type" id={'typeselect-' + type.type}
           checked={type.type === selectedType}
           value={type.type}
           on:click={() => selectType(type.type)}>
    <label for={'typeselect-' + type.type}>{type.label}</label>
  {/each}
</div>

<style>
</style>
