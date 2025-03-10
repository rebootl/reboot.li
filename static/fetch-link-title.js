/** @typedef {HTMLElement} HTMLElement
  * @description This is a custom element that fetches the title of a webpage given its URL.
  * @example
  * <fetch-link-title>
  *   <input type="url" name="url" placeholder="URL" required>
  *   <span title-status></span>
  *   <input type="text" name="title" placeholder="Title" required>
  * </fetch-link-title>
  */
class FetchLinkTitle extends HTMLElement {

  constructor() {
    super();
    this.timeoutId = null;
  }

  connectedCallback() {
    /** @type {HTMLInputElement} */
    this.urlInput = this.querySelector('input[name="url"]');
    /** @type {HTMLInputElement} */
    this.titleInput = this.querySelector('input[name="title"]');
    /** @type {HTMLSpanElement} */
    this.titleStatus = this.querySelector('[title-status]');
    /** Array of required elements */
    const elements = [
      this.urlInput,
      this.titleInput,
      this.titleStatus,
    ];
    if (elements.some((element) => !element)) {
      console.error('Required element not found.');
      return;
    }

    this.urlInput.addEventListener('input', () => {
      this.debounce(() => this.fetchTitle(), 500);
    });
  }

  /**
   * Debounce a function
   * @param {Function} func Function to be debounced
   * @param {number} wait Time to wait before function is called
   */
  debounce(func, wait) {
    if (this.timeoutId) {
      clearTimeout(this.timeoutId);
    }
    this.timeoutId = setTimeout(func, wait);
  }

  /**
   * Fetch title of the webpage given its URL
   */
  async fetchTitle() {
    /** @type {string} */
    const url = this.urlInput.value.trim();
    if (!url) return;

    this.titleStatus.textContent = 'Fetching title...';
    try {
      const response = await fetch(`/api/get-title?url=${encodeURIComponent(url)}`);
      if (response.status === 200) {
        const data = await response.json();
        if (data.success) {
          this.titleStatus.textContent = 'Title fetched.';
          this.titleInput.value = data.title;
        } else {
          this.titleStatus.textContent = `Error getting title: ${data.error}`;
        }
      } else {
        this.titleStatus.textContent = `${response.status}: ${response.statusText}`;
      }
    } catch (error) {
      this.titleStatus.textContent = `Error: ${error.message}`;
      console.error('Error:', error);
    }
  }
}

customElements.define('fetch-link-title', FetchLinkTitle);
