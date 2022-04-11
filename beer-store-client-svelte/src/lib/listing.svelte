<script lang="ts">
	import Item from './item.svelte';

	let items = [{id:1,name:"aaa"},{id:2,name:"bbb"},{id:3,name:"ccc"},{id:4,name:"ddd"}];
</script>

{#each items as item (item.id) }
  <Item beer={item}/>
{/each}