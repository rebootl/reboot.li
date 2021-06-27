<svelte:head>
	<title>betterreflect-app</title>
</svelte:head>

<script context="module">
	/*
	export const load = async ({ fetch }) => {
		const res = await fetch('/users.json');

		if (res.ok) {
			console.log(res)
			const users = await res.json();

			return {
				props: { users }
			};
		}

		const { message } = await res.json();

		return {
			error: new Error(message)
		};
	};*/
</script>

<script>
	//import getDb from '$lib/db';
	import Login from '$lib/Login.svelte';
	import { onMount } from 'svelte';

  let users = [];

  async function getUsers() {
		try {
			const res = await fetch('/users.json');
			if (res.ok) {
				users = await res.json();
				return;
			}
			const { message } = await res.json();
	    console.log(new Error(message));
		} catch(error) {
			console.error(error);
		}
	}

	onMount(async () => {
		getUsers();
	});
</script>

<h1>betterreflect-app</h1>

<h2>list of users</h2>

{#each users as u}
 {u.username}<br>
{/each}

<Login />

<style></style>
