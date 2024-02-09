<script>
  import { browser } from '$app/environment';
  import { onDestroy, onMount } from 'svelte';

	/** @typedef {Object} Props
   * @property {import('$lib/server/db.js').ImageData[]} images
   * @property {number} currentImageIdx
   * @property {boolean} show
   * @property {() => void} close
   */
  /** @param {Props} props */
	let { show, images, currentImageIdx,
		close } = $props();

	/** @type {HTMLButtonElement|null} */
	let closeButton = $state(null);
	
	function nextImage() {
		if (currentImageIdx < images.length - 1) {
	    currentImageIdx++;
	  }
  }

	function previousImage() {
	  if (currentImageIdx > 0) {
      currentImageIdx--;
    }
  }

	/** @param {number} i */
	/*function setImage(i) {
	  currentImageIdx = i;
  }*/

	/** @param {KeyboardEvent} e */
  function keydownHandler(e) {
    if (e.key === 'ArrowRight') {
      nextImage();
    } else if (e.key === 'ArrowLeft') {
      previousImage();
    }
  }

	if (browser) {
    window.addEventListener('keydown', keydownHandler);
  }

	onMount(() => {
    closeButton?.focus();
		document.querySelector('body')?.setAttribute('style', 'overflow: hidden');
  });

  onDestroy(() => {
    window.removeEventListener('keydown', keydownHandler);
    document.querySelector('body')?.removeAttribute('style');
  });
</script>

<div id="image-viewer" class:show={show}
     role="dialog"
     aria-label="Image Viewer"
 		 aria-hidden={!show}
     aria-modal="true"
     >
	<div class="image-box">
		<div class="image-header">
			<button class="icon-button close-button" onclick={close} bind:this={closeButton}
				      aria-label="Close Image Viewer"
							>
				<span class="material-icons">close</span>
	    </button>
    </div>
	  <img class="current-image" src={ images[currentImageIdx].path }
				 alt={ images[currentImageIdx].comment } />
		<div class="image-footer">
			<button class="icon-button" onclick={() => previousImage()}
							aria-label="Previous Image"
							>
				<span class="material-icons">navigate_before</span>
      </button>
			<div class="image-footer-inner">
				<div class="image-index-indicator">
					{#each images as image, i}
	            {#if i === currentImageIdx}
	              <span class="material-icons image-index-dot">lens</span>
	            {:else}
	              <span class="material-icons-outlined image-index-dot"
											>lens</span>
	              <!--<span class="material-icons-outlined image-index-dot"
											onclick={() => setImage(i)}
											onkeydown={() => setImage(i)}
											role="button"
                      tabindex="0"
                      aria-label={ `Show image ${i + 1}` }
											>lens</span>-->
	            {/if}
		      {/each}
	      </div>
			  <div class="comment">
					<small>{ images[currentImageIdx].comment }</small>
				</div>
			</div>
			<button class="icon-button" onclick={() => nextImage()}
              aria-label="Next Image"
              >
        <span class="material-icons">navigate_next</span>
      </button>
		</div>
	</div>
</div>

<style>
  #image-viewer {
    position: fixed;
    top: 0;
    left: 0;
    z-index: 100;
    width: 100vw;
    height: 100vh;
    background-color: rgba(0,0,0,0.7);
    display: none;
		flex-direction: column;
  }
  #image-viewer.show {
    display: flex;
  }
  .image-box {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    /*height: calc(100vh - var(--header-height));*/
    height: 100%;
  }
  .image-header {
		height: var(--header-height);
	  width: 100%;
		position: relative;
	}
	.close-button {
	  position: absolute;
	  bottom: 2px;
	  right: 0;
		height: calc(var(--header-height) - 2px);
	}
  .current-image {
		border-radius: 0;
	 	max-height: calc(100vh - 2 * var(--header-height));
		max-width: calc(100vw);
  }
  .image-footer {
    display: flex;
    justify-content: space-between;
		height: var(--header-height);
    width: 100%;
  }
  .image-footer-inner {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }
  .image-index-indicator {
    display: flex;
		flex-direction: row;
		margin-top: 5px;
  }
  .image-index-dot {
		width: 20px;
		font-size: 10px;
  }
  .comment {
		display: flex;
	  justify-content: center;
	  align-items: center;
		width: 100%;
		/*background-color: rgba(0,0,0,0.7);*/
		color: var(--text-color-dimmed);
		height: 100%;
  }
	.icon-button {
		background-color: rgba(0,0,0,0.7);
  }
</style>
