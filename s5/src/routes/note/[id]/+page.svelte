<script>
	import { marked } from 'marked';

	export let data;
  // console.log(data);
  
	const entry = data.entry;

  const text = marked.parse(entry.content);
</script>

<div class="note-header">
	<div>
		<small>{entry.created_at}</small>
		{#if entry.private}
			<small><span class="material-icons">lock</span> Private</small>
		{/if}
  </div>
	{#if data.clientData.loggedIn}
		<small><a href="/editNote/{entry.id}"><span class="material-icons">edit</span></a></small>
	{/if}
</div>

<h1>{entry.title}</h1>

{@html text}

<style>
	.note-header {
    display: flex;
    justify-content: space-between;
		padding-top: 10px;
  }
  .note-header small {
    /*font-size: 0.85em;*/
    color: var(--text-color-dimmed);
  }
  .note-header small .material-icons {
    font-size: small;
  }
</style>
