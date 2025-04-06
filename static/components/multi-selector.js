/** @typedef {HTMLElement} HTMLElement
  * @description This is a custom element that provides a multi-select feature.
  * The user can select an item from a list and add it to a list of selected items.
*/
class MultiSelector extends HTMLElement {

  constructor() {
    super();
  }

  connectedCallback() {
    /** @type {HTMLSelectElement} */
    this.baseSelectElement = this.querySelector('.multi-select');
    /** @type {HTMLSelectElement} */
    this.itemSelectElement = this.querySelector('.item-selector');
    /** @type {HTMLButtonElement} */
    this.addButton = this.querySelector('.add-button');
    /** @type {HTMLUListElement} */
    this.listElement = this.querySelector('.selected-items');
    /** @type {HTMLCollectionOf<HTMLButtonElement>} */
    this.removeButtons = this.querySelectorAll('.remove-button');
    /** @type {HTMLTemplateElement} */
    this.selectedItemTemplate = document.querySelector('.selected-item-template');

    /** Array of required elements */
    const elements = [
      this.baseSelectElement,
      this.itemSelectElement,
      this.addButton,
      this.listElement,
      this.selectedItemTemplate
    ];
    if (elements.some((element) => !element)) {
      console.error('No required element found');
      return;
    }

    // NOTE: the use of the arrow function here preserves the 'this' context
    // otherwise this.addItem.bind(this)) has to be used instead, this is
    // the case for _callbacks_
    this.addButton.addEventListener('click', (e) => this.addItem(e));
    this.removeButtons.forEach((button) => {
      button.addEventListener('click', (e) => this.removeItem(e));
    });
  }

  /**
   * @param {Event} e
   */
  addItem(e) {
      e.preventDefault();
      const selectedOption = this.itemSelectElement.options[this.itemSelectElement.selectedIndex];
      if (selectedOption) {
        const value = selectedOption.value;
        const existingOption = this.baseSelectElement.querySelector(`option[value="${value}"]`);
        if (!existingOption || !existingOption.selected) {
          existingOption.setAttribute('selected', '');
          this.addListElement(value);
        }
      }
  }

  /**
   * @param {Event} e
   */
  removeItem(e) {
    e.preventDefault();
    const itemName = e.target.getAttribute('data-item-name');
    const listItem = e.target.parentNode;
    this.listElement.removeChild(listItem);

    const option = this.baseSelectElement.querySelector(`option[value="${itemName}"]`);
    option.removeAttribute('selected');
  }

  /**
   * @param {string} text
   */
  addListElement(text) {
    const newListitem = this.selectedItemTemplate.content.cloneNode(true);
    const span = newListitem.querySelector('span');
    span.textContent = text;

    const removeButton = newListitem.querySelector('.remove-button');
    removeButton.setAttribute('data-item-name', text);
    removeButton.addEventListener('click', (e) => this.removeItem(e));
    this.listElement.appendChild(newListitem);
  }
}

customElements.define('multi-selector', MultiSelector);
