<script lang="ts">
	import Item from './item.svelte';

	import type { Beer } from './Beer';

	export let beers: Array<Beer>;
</script>

{#each beers as beer (beer.idbeer)}
	<Item {beer} />
{/each}
