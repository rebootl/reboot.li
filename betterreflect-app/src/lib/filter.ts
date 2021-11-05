export function getFilteredEntries(entries, v) {
  const [ selectedTopics, selectedTags ] = v;
  let f = entries;
  if (selectedTopics.length > 0) {
    f = f.filter(e => {
      for (const t of e.topics) {
        if (selectedTopics.includes(t)) return e;
      }
    });
  }
  if (selectedTags.length > 0) {
    f = f.filter(e => {
      for (const t of e.tags) {
        if (selectedTags.includes(t)) return e;
      }
    });
  }
  return f;
}
