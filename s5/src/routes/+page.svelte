<script>
	// import SideNav from '$lib/SideNav.svelte';

	import { marked } from 'marked';

  const text = `# Welcome to reboot.li

This is the 1'938'489'483th iteration of my personal website.
Why it is like this, I don't know. Well, maybe because I use this
as a playground for experimentation and trying out new stuff.

My name is Cem. I'm a web dev, kinda, and sysadmin, maybe?

![me](/me-small.jpeg)

I'm 40+, yikes.

My pronouns are _he/him_.

## Links

* [Github](https://github.com/rebootl)
* [Twitch](https://www.twitch.tv/rebootl)
* [Lichess](https://lichess.org/@/rebootl)

`
  const html = marked.parse(text);
</script>

{@html html}
