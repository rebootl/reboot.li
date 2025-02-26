class MultiSelector extends HTMLElement {

  constructor() {
    super();
  }

  connectedCallback() {
    this.baseSelectElement = this.querySelector('.multi-select');
    this.itemSelectElement = this.querySelector('.item-selector');
    this.addButton = this.querySelector('.add-button');
    this.listElement = this.querySelector('.selected-items');
    this.removeButtons = this.querySelectorAll('.remove-button');

    // NOTE: the use of the arrow function here preserves the 'this' context
    // otherwise this.addItem.bind(this)) has to be used instead, this is
    // the case for _callbacks_
    this.addButton.addEventListener('click', (e) => this.addItem(e));
    this.removeButtons.forEach((button) => {
      button.addEventListener('click', (e) => this.removeItem(e));
    });
  }

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

  removeItem(e) {
    e.preventDefault();
    const itemName = e.target.getAttribute('data-item-name');
    const listItem = e.target.parentNode;
    this.listElement.removeChild(listItem);

    const option = this.baseSelectElement.querySelector(`option[value="${itemName}"]`);
    option.removeAttribute('selected');
  }

  addListElement(text) {
    const newListitem = document.querySelector('.selected-item').cloneNode(true);
    const span = newListitem.querySelector('span');
    span.textContent = text;

    const removeButton = newListitem.querySelector('.remove-button');
    removeButton.setAttribute('data-item-name', text);
    removeButton.addEventListener('click', (e) => this.removeItem(e));
    this.listElement.appendChild(newListitem);
  }
}

customElements.define('multi-selector', MultiSelector);
