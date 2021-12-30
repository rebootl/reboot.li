let timeoutID;
export function debounce(f, t) {
  if (timeoutID) {
    clearTimeout(timeoutID);
    timeoutID = setTimeout(f, t);
  } else {
    timeoutID = setTimeout(f, t);
  }
}
