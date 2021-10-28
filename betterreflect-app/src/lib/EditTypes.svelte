<script>
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let selectedType = 'any';

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
    }
  ];

  function selectType(t) {
    selectedType = t;
    dispatch('change', t);
  }
</script>

<div class="types">
  Type:
  {#each types as type}
    <input type="radio" name="edittype" id={'edittypeselect-' + type.type}
           checked={type.type === selectedType}
           value={type.type}
           on:click={() => selectType(type.type)}>
    <label for={'edittypeselect-' + type.type}>{type.label}</label>
  {/each}
</div>

<style>
</style>
