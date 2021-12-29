<script context="module">
	import SideNav from '$lib/SideNav.svelte';

	import { marked } from 'marked';


	export const prerender = true;

	const text = `# About

## this

This is my personal website. A place for me to __share and organize__ links,
tasks, images and more.

Also it's an ongoing web development project to learn and try out new things.

Content is my own, unless otherwise stated.

## me

My name is Cem. Online I use the alias _rebootl_.

As of 2021, I'm 40 years old.

Amongst other things I like computers, chess and astronomy.

My code is on github: https://github.com/rebootl

Occasionally I stream on twitch: https://www.twitch.tv/rebootl

My lichess account: https://lichess.org/@/rebootl

I use discord: rebootl#6863 You can find me for example on the lichess server.

My email is: cem.aydin at gmx.ch

### Software I use

__OS:__ Arch Linux / Kubuntu / Debian (Servers)<br>
I've been using Linux exclusively for years now and I reallly love it.<br>
__Desktop:__ KDE Plasma, i3, openbox<br>
__Browser:__ firefox<br>
__Streaming:__ obs<br>
__Editors:__ atom, nano<br>

`;

  const html = marked.parse(text);

</script>

<SideNav entries={[]} showOnWide={false} />
<main>
	{@html html}
</main>

<style>
	.copy {
		color: blue;
	}
</style>
