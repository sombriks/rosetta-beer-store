<script lang="ts">
	import type { Beer } from './Beer';

	export let beer: Beer;
</script>

<div class="item">
	<h3>{beer.titlebeer}</h3>
	<p>{beer.descriptionbeer}</p>
</div>


<style>
.item {
	display:flex;
	justify-items: center;
	align-items: center;
	justify-content: space-between;
}
</style>