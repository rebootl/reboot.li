import { writable } from 'svelte/store';


export const showMenu = writable(false);

export const currentTopics = writable([]);
export const currentTags = writable([]);

export const currentTagsByTopics = writable({});
