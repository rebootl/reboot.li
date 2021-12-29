<script context="module">
	export const prerender = true;
</script>

<script>
	import SideNav from '$lib/SideNav.svelte';

	import { marked } from 'marked';

	const text = `# About

## this

This is my personal website. A place for me to share and organize links,
tasks, images and more.

Also it's an ongoing web development project to learn and try out new things.

## me

![me](/me-small.jpeg)

My name is Cem. My pronouns are _he/him_. Online I use the alias _rebootl_.

As of 2021, I'm 40 years old. Amongst other things I like computers, chess and
astronomy.

My code is on __github__: https://github.com/rebootl

Occasionally I stream on __twitch__: https://www.twitch.tv/rebootl

My __lichess__ account: https://lichess.org/@/rebootl

I use __discord__: _rebootl#6863_<br>
You can find me for example on the lichess server.

Additionally I'm on libera __IRC__ as rebootl. Channel: #revamp

I also have telegram.

### Software I use

__OS:__ Arch Linux / Kubuntu / Debian (Servers)<br>
I've been using Linux exclusively for years now and I reallly love it.<br>
__Desktop:__ KDE Plasma, i3, openbox<br>
__Browser:__ firefox<br>
__Streaming:__ obs<br>
__Editors:__ atom, nano<br>

## copying

Unless otherwise stated content is my own and copyright applies.

Copyright 2021 Cem Aydin
`;

  const html = marked.parse(text);

</script>

<SideNav entries={[]} showOnWide={false} />
<main>
	{@html html}
</main>
