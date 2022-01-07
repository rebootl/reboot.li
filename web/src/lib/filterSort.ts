export function getFilteredEntries(entries, selectedTopics, selectedTags,
  filterText) {

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
  if (filterText !== '') {
    const lc = filterText.toLowerCase();
    f = f.filter(e => {
      for (const t of e.topics) {
        if (t.toLowerCase().includes(lc)) return e;
      }
      for (const t of e.tags) {
        if (t.toLowerCase().includes(lc)) return e;
      }
      if (e.type !== 'image') {
        if (e.text.toLowerCase().includes(lc)) return e;
      } else {
        for (const i of e.images) {
          if (i.comment.toLowerCase().includes(lc)) return e;
        }
      }
    });
  }
  const p = f.filter(e => e.pinned).sort(sortByDateNewestFirst);
  const q = f.filter(e => !e.pinned).sort(sortByDateNewestFirst);
  return [ ...p, ...q ];
}

function sortByDate(a, b) {
  return new Date(a.date) - new Date(b.date);
}

function sortByDateNewestFirst(a, b) {
  return new Date(b.date) - new Date(a.date);
}
