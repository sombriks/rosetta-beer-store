<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	import type { Beer } from './Beer';
  
	const dispatch = createEventDispatcher();
  
	let search = '';
	let page = 1;
	let pageSize = 10;
	
  export let beers: Array<Beer>;

  function onSearch(p = 0){
    page += p
    dispatch("search",{
      search, page, pageSize
    })
  }

</script>

<div class="searchbox">
	<input bind:value={search} />
	<button on:click="{e => onSearch()}">Search</button>
	<button disabled={page <= 1} on:click="{e => onSearch(-1)}">Prev</button>
	<button disabled={beers.length < 10} on:click="{e => onSearch(1)}">Next</button>
</div>

<style>
	.searchbox {
		display: flex;
	}
	.searchbox > input {
		flex-grow: 7;
	}
</style>
