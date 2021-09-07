<script>
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  let types = [
    {
      icon: 'task_alt',
      label: 'Task',
			type: 'task'
    },
    {
      icon: 'link',
      label: 'Link',
			type: 'link'
    },
    {
      icon: 'article',
      label: 'Article',
			type: 'article'
    },
    {
      icon: 'image',
      label: 'Image',
			type: 'image'
    }
  ];

  let selectedType = 'any';

  function selectType(t) {
    selectedType = t;
    console.log(t)
    dispatch('change', t);
  }
</script>

<div class="types">
  Type:
  {#each types as type}
    <input type="radio" name="edittype" id={'edittypeselect-' + type.type}
           checked={type.type === 'task'}
           value={type.type}
           on:click={() => selectType(type.type)}>
    <label for={'edittypeselect-' + type.type}>{type.label}</label>
  {/each}
</div>

<style>
</style>
