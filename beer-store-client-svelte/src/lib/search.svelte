<script lang="ts">
	let search = '';
</script>

<div class="searchbox">
  <input bind:value={search}/>
  <button>Search</button>
  <button>Prev</button>
  <button>Next</button>
</div>

<style>
  .searchbox {
    display: flex;

  }
  .searchbox > input {
    flex-grow: 7;
  }
</style>