/** @typedef {HTMLElement} HTMLElement
  * @description This is a custom element that will confirm the submission of a form.
*/
class ConfirmSubmit extends HTMLElement {

  constructor() {
    super();
  }

  connectedCallback() {
    /** @type {HTMLButtonElement} */
    this.submitButton = this.querySelector('button');
    if (!this.submitButton) {
      console.error('No button found');
      return;
    }
    /** @type {string} */
    this.message = this.getAttribute('message') || 'Are you sure you want to submit?';
    /** @type {boolean} */
    this.replaceAction = this.hasAttribute('replaceAction');
    /** @type {string} */
    this.action = this.getAttribute('replaceAction') || '';

    // NOTE: the use of the arrow function here preserves the 'this' context
    // otherwise this.addItem.bind(this)) has to be used instead, this is
    // the case for _callbacks_
    this.submitButton.addEventListener('click', (e) => this.confirmSubmit(e));
  }

  /**
   * @param {MouseEvent} e
   */
  confirmSubmit(e) {
    e.preventDefault();
    if (confirm(this.message)) {
      if (this.replaceAction) {
        this.submitButton.form.action = this.action;
      }
      this.submitButton.form.submit();
    }
  }
}

customElements.define('confirm-submit', ConfirmSubmit);
