<script>
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

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

  let selectedType = 'any';

  function selectType(t) {
    selectedType = t;
    console.log(t)
    dispatch('change', t);
  }

//           checked={type.type === 'all'}
</script>

<div class="types">
  Types:
  {#each types as type}
    <input type="radio" name="type" id={'typeselect-' + type.type}
           checked={type.type === 'any'}
           value={type.type}
           on:click={() => selectType(type.type)}>
    <label for={'typeselect-' + type.type}>{type.label}</label>
  {/each}
</div>

<style>
</style>
