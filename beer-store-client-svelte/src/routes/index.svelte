<script lang="ts">
	import { onMount } from 'svelte';

	import { listBeers } from '$lib/api';

	import Listing from '$lib/listing.svelte';
	import Search from '$lib/search.svelte';

	let beers = [];

	async function doSearch(e) {
		beers = await listBeers(e.detail);
	}

	onMount(() => {
		doSearch({ detail: { search: '', page: 1, pageSize: 10 } });
	});
</script>

<h1>Beer store</h1>
<Search {beers} on:search={doSearch} />
<Listing {beers} />
