/** @typedef {HTMLElement} HTMLElement
  * @description This is a custom element that will read a message of the day from a text file
  * and display it in the element.
*/
class MessageOfTheDay extends HTMLElement {

  constructor() {
    super();
  }

  connectedCallback() {
    /** @type {HTMLElement} */
    this.containerElement = this.querySelector('pre');
    // NOTE: when the web component is rendered through markdown it cannot find the inner element
    //       therefore we create it inside the check below
    //       it's not clear why it cannot find the element because when looking at the site source
    //       it is actually there
    if (!this.containerElement) {
      // console.warn('No container element found for message of the day, creating one.');
      this.containerElement = document.createElement('pre');
      this.containerElement.classList.add('motd');
      this.appendChild(this.containerElement);
    }
    /** @type {string} */
    this.path = this.getAttribute('path') || '/static/motd.txt';

    this.setMessage();
  }

  async setMessage() {
    // check if the path is set
    if (!this.path) {
      console.error('Path not set for message of the day');
      return;
    }
    // fetch the message from the text file
    try {
      const response = await fetch(this.path);
      if (response.ok) {
        const text = await response.text();
        this.containerElement.innerHTML = text;
      } else {
        console.error(`Failed to fetch message of the day: ${response.statusText}`);
      }
    } catch (error) {
      console.error(`Error fetching message of the day: ${error}`);
    }
  }
}

customElements.define('message-of-the-day', MessageOfTheDay);
